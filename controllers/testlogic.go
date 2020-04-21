package controllers

import "encoding/json"

type T_In26 struct {
	Param_1 int
}

type T_Out26 struct {
	Param_2 string
}
type Panduanceshi20041301_26 struct {
	NCController
}

// @router /Panduanceshi20041301_26 [Post]
func (c *Panduanceshi20041301_26) Panduanceshi20041301_26() {
	var In T_In26
	var Out T_Out26
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &In)
	if c.handlerErrOK(err) {
		c.LogicBody(&In, &Out)
		c.responseSuccess(map[string]interface{}{"output": Out})
	}
}
func (c *Panduanceshi20041301_26) LogicBody(In *T_In26, Out *T_Out26) {

	if In.Param_1 == 1 {
		Out.Param_2 = `类型1`
	}
	if In.Param_1 == 2 {
		Out.Param_2 = `类型2`
	}
	if ((In.Param_1 > 2) && (In.Param_1 < 5)) || ((In.Param_1 > 12) && (In.Param_1 < 15)) {
		Out.Param_2 = `类型3`
	}

	//node1B_1:=len("Node1.A")//左侧
	//node1B_2:=(node1B_1+12)//左侧
	//node1B_3:=(node1B_2*32)//左侧
	//node1B_4:=len("Node1.A")//左侧
	//node1B_5:=(node1B_4+12)//左侧
	//node1B_6:=node1B_5*32//全部
	//Node1.B=node1B_6

}
