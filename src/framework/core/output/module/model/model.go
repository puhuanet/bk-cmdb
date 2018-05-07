/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except 
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and 
 * limitations under the License.
 */
 
package model

import (
	"configcenter/src/framework/common"
	"configcenter/src/framework/core/output/module/client"
	"configcenter/src/framework/core/types"
)

var _ Model = (*model)(nil)

// model the metadata structure definition of the model
type model struct {
	ObjCls      string `field:"bk_classification_id"`
	ObjIcon     string `field:"bk_obj_icon"`
	ObjectID    string `field:"bk_obj_id"`
	ObjectName  string `field:"bk_obj_name"`
	IsPre       bool   `field:"ispre"`
	IsPaused    bool   `field:"bk_ispaused"`
	Position    string `field:"position"`
	OwnerID     string `field:"bk_supplier_account"`
	Description string `field:"description"`
	Creator     string `field:"creator"`
	Modifier    string `field:"modifier"`
}

func (cli *model) ToMapStr() types.MapStr {
	return types.MapStr{
		SupplierAccount:  cli.OwnerID,
		ObjectID:         cli.ObjectID,
		ObjectIcon:       cli.ObjIcon,
		ClassificationID: cli.ObjCls,
		ObjectName:       cli.ObjectName,
		IsPre:            cli.IsPre,
		IsPaused:         cli.IsPaused,
		Position:         cli.Position,
		Description:      cli.Description,
		Creator:          cli.Creator,
		Modifier:         cli.Modifier,
	}
}

func (cli *model) Attributes() ([]Attribute, error) {

	cond := common.CreateCondition().Field(ObjectID).Like(cli.ObjectID).Field(SupplierAccount).Eq(cli.OwnerID)

	dataMap, err := client.GetClient().CCV3().Attribute().SearchObjectAttributes(cond)

	if nil != err {
		return nil, err
	}

	attrs := make([]Attribute, 0)

	for _, item := range dataMap {
		tmpItem := &attribute{}
		common.SetValueToStructByTags(tmpItem, item)
		attrs = append(attrs, tmpItem)
	}

	return attrs, nil

}

func (cli *model) Save() error {

	// construct the search condition
	cond := common.CreateCondition().Field(ObjectID).Eq(cli.ObjectID).Field(ObjectName).Eq(cli.ObjectName)

	// search all objects by condition
	dataItems, err := client.GetClient().CCV3().Model().SearchObjects(cond)
	if nil != err {
		return err
	}

	// create a new object
	if 0 == len(dataItems) {
		if _, err = client.GetClient().CCV3().Model().CreateObject(cli.ToMapStr()); nil != err {
			return err
		}
		return nil
	}

	// update the exists one
	for _, item := range dataItems {

		item.Set(ObjectIcon, cli.ObjIcon)
		item.Set(ClassificationID, cli.ObjCls)
		item.Set(ObjectName, cli.ObjectName)
		item.Set(IsPre, cli.IsPre)
		item.Set(IsPaused, cli.IsPaused)
		item.Set(Position, cli.Position)
		item.Set(Description, cli.Description)
		item.Set(Modifier, cli.Modifier)

		item.Remove(ObjectID)

		id, err := item.Int("id")
		if nil != err {
			return err
		}

		cond := common.CreateCondition()
		cond.Field(ObjectID).Eq(cli.ObjectID).Field(SupplierAccount).Eq(cli.OwnerID).Field("id").Eq(id)
		if err = client.GetClient().CCV3().Model().UpdateObject(item, cond); nil != err {
			return err
		}
	}

	// success
	return nil
}

func (cli *model) CreateAttribute() Attribute {
	attr := &attribute{
		ObjectID:      cli.ObjectID,
		OwnerID:       cli.OwnerID,
		Creator:       cli.Creator,
		PropertyGroup: "default",
	}
	return attr
}

func (cli *model) SetClassification(classificationID string) {
	cli.ObjCls = classificationID
}

func (cli *model) GetClassification() string {
	return cli.ObjCls
}

func (cli *model) SetIcon(iconName string) {
	cli.ObjIcon = iconName
}

func (cli *model) GetIcon() string {
	return cli.ObjIcon
}

func (cli *model) SetID(id string) {
	cli.ObjectID = id
}

func (cli *model) GetID() string {
	return cli.ObjectID
}

func (cli *model) SetName(name string) {
	cli.ObjectName = name
}
func (cli *model) GetName() string {
	return cli.ObjectName
}

func (cli *model) SetPaused() {
	cli.IsPaused = true
}

func (cli *model) SetNonPaused() {
	cli.IsPaused = false
}

func (cli *model) Paused() bool {
	return cli.IsPaused
}

func (cli *model) SetPosition(position string) {
	cli.Position = position
}

func (cli *model) GetPosition() string {
	return cli.Position
}

func (cli *model) SetSupplierAccount(ownerID string) {
	cli.OwnerID = ownerID
}
func (cli *model) GetSupplierAccount() string {
	return cli.OwnerID
}

func (cli *model) SetDescription(desc string) {
	cli.Description = desc
}
func (cli *model) GetDescription() string {
	return cli.Description
}
func (cli *model) SetCreator(creator string) {
	cli.Creator = creator
}
func (cli *model) GetCreator() string {
	return cli.Creator
}
func (cli *model) SetModifier(modifier string) {
	cli.Modifier = modifier
}
func (cli *model) GetModifier() string {
	return cli.Modifier
}
func (cli *model) CreateGroup() Group {
	g := &group{
		OwnerID:  cli.OwnerID,
		ObjectID: cli.ObjectID,
	}
	return g
}

func (cli *model) FindAttributesLikeName(attributeName string) (AttributeIterator, error) {
	cond := common.CreateCondition().Field(PropertyName).Like(attributeName)
	return newAttributeIterator(cond)
}
func (cli *model) FindAttributesByCondition(cond common.Condition) (AttributeIterator, error) {
	return newAttributeIterator(cond)
}
func (cli *model) FindGroupsLikeName(groupName string) (GroupIterator, error) {
	cond := common.CreateCondition().Field(GroupName).Like(groupName)
	return newGroupIterator(cond)
}
func (cli *model) FindGroupsByCondition(cond common.Condition) (GroupIterator, error) {
	return newGroupIterator(cond)
}
