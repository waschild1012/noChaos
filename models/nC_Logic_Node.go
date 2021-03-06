// @Title  nC_Logic_Node
// @Description  节点的整体处理
package models

import (
	"com.waschild/noChaos-Server/utils"
	"fmt"
	"strings"
)

//操作节点
type NC_Node struct {
	Name     string      //节点名称
	Mark     string      //节点地址
	Type     NodeType    //节点类型
	Declare  interface{} //声明内容
	InLines  []NC_Flow   //入线
	OutLines []NC_Flow   //出线
}

type NodeType string

const (
	AssignNode   NodeType = "assign"   //赋值节点
	VariableNode NodeType = "variable" //定义节点
	CycleNode    NodeType = "cycle"    //循环节点
	JudgeNode    NodeType = "judge"    //判断节点
	LogicNode    NodeType = "logic"    //逻辑节点
	FormNode     NodeType = "form"     //表单处理节点
)

//定义节点
type NC_VariableNode struct {
	Mark       string        //地址编号
	Properties []NC_Property //类型结构属性
}

//单个定义的变量
type NC_Property struct {
	Name         string        //名称
	Mark         string        //标识
	Type         string        //类型
	Multiple     bool          //多个
	InitialValue string        //初始值
	Properties   []NC_Property //类型结构属性
}

// TODO NC_Node-获取赋值节点源码assign
func (node *NC_Node) getAssignCode() string {
	var split []string
	assigns := []Assign{}
	utils.ReUnmarshal(node.Declare, &assigns)
	for _, assign := range assigns {
		split = append(split, assign.AnalyzeExpression())
	}
	return strings.Join(split, "\t\n") + "\t\n"
}

// TODO NC_Node-获取定义节点源码variable
func (node *NC_Node) getVariableCode() string {
	code := "var " + node.Mark + " " + getStructName(node.Mark) + "\n"
	variables := []NC_Property{}
	utils.ReUnmarshal(node.Declare, &variables)
	for _, p := range variables {
		if p.InitialValue != "" {
			code += node.Mark + "." + p.Mark + "=" + p.InitialValue + "\n"
		}
	}
	return code
}

// TODO NC_Node-获取循环节点源码cycle
func (node *NC_Node) getCycleCode(flow NC_Flow) string {
	judgeCode := fmt.Sprintf("if %s == false {\n\tbreak\n}\n", flow.getJudgeCode())
	return "for { \n" + judgeCode
}

// TODO NC_Node-获取判断节点源码judge
func (node *NC_Node) getJudgeCode(flow NC_Flow) string {
	return fmt.Sprintf("if %s {", flow.getJudgeCode())
}

// TODO NC_Node-获取逻辑节点源码logic
func (node *NC_Node) getLogicCode() (string, string) {
	logic_node := Logic_Node{}
	utils.ReUnmarshal(node.Declare, &logic_node)
	return logic_node.GetCode(node.Mark), logic_node.Package
}

// TODO NC_Node-获取表单节点源码form
func (node *NC_Node) getFormCode() string {
	code := ""
	return code
}
