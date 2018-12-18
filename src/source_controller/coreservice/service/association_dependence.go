/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.,
 * Copyright (C) 2017,-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the ",License",); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an ",AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"configcenter/src/source_controller/coreservice/core"
	"configcenter/src/source_controller/coreservice/core/association"
	"configcenter/src/source_controller/coreservice/core/instances"
	"configcenter/src/storage/dal"
)

type associationDepend struct {
	instanceOperation core.InstanceOperation
}

func NewAssociationDepend(dbProxy dal.RDB) association.OperationDependences {
	asstDepend := &associationDepend{}
	asstDepend.instanceOperation = instances.New(dbProxy)
	return asstDepend

}
func (m *associationDepend) IsInstanceExist(ctx core.ContextParams, objID string, instID uint64) (exists bool, err error) {
	return true, nil
}
