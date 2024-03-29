package main

import (
	"apiserver/config"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// init config
	if err := config.Init(*pflag.StringP("config", "c", "", "gormgenrate config file path.")); err != nil {
		panic(err)
	}
	mySQLDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"),
		true,
		"Local")
	// 连接数据库
	db, err := gorm.Open(mysql.Open(mySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	// 生成实例
	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath:      "../../dal/dockerdb/dockerquery",
		ModelPkgPath: "dockermodel",
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable
		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values
		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: true, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
	})
	// 设置目标 db
	g.UseDB(db)

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(detailType string) (dataType string){
		"int": func(detailType string) (dataType string) {
			return "int64"
		},
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 指定json标签命名策略
	jsonTagNameStrategy := func(columnName string) (tagContent string) {
		return strutil.CamelCase(columnName)
	}
	// 要先于`ApplyBasic`执行
	g.WithJSONTagNameStrategy(jsonTagNameStrategy)

	// 创建模型文件
	User := g.GenerateModel("tb_user")
	UserToken := g.GenerateModel("tb_user_token",
		gen.FieldJSONTag("user_id", "userID"))
	Article := g.GenerateModel("tb_articles",
		gen.FieldNew("Editor", "string", `gorm:"-"`),
		gen.FieldType("uid", "uint64"),
		gen.FieldType("cate_id", "uint64"),
		gen.FieldJSONTag("cate_id", "cateID"),
	)

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic(User, UserToken, Article)
	g.Execute()
}
