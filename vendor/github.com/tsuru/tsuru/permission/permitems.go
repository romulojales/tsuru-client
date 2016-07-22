// AUTOMATICALLY GENERATED FILE - DO NOT EDIT!
// Please run 'go generate' to update this file.
//
// Copyright 2016 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package permission

var (
	PermAll                              = PermissionRegistry.get("")
	PermApp                              = PermissionRegistry.get("app")
	PermAppAdmin                         = PermissionRegistry.get("app.admin")
	PermAppAdminQuota                    = PermissionRegistry.get("app.admin.quota")
	PermAppAdminRoutes                   = PermissionRegistry.get("app.admin.routes")
	PermAppAdminUnlock                   = PermissionRegistry.get("app.admin.unlock")
	PermAppCreate                        = PermissionRegistry.get("app.create")
	PermAppDelete                        = PermissionRegistry.get("app.delete")
	PermAppDeploy                        = PermissionRegistry.get("app.deploy")
	PermAppDeployArchiveUrl              = PermissionRegistry.get("app.deploy.archive-url")
	PermAppDeployBuild                   = PermissionRegistry.get("app.deploy.build")
	PermAppDeployGit                     = PermissionRegistry.get("app.deploy.git")
	PermAppDeployImage                   = PermissionRegistry.get("app.deploy.image")
	PermAppDeployRollback                = PermissionRegistry.get("app.deploy.rollback")
	PermAppDeployUpload                  = PermissionRegistry.get("app.deploy.upload")
	PermAppRead                          = PermissionRegistry.get("app.read")
	PermAppReadDeploy                    = PermissionRegistry.get("app.read.deploy")
	PermAppReadEnv                       = PermissionRegistry.get("app.read.env")
	PermAppReadEvents                    = PermissionRegistry.get("app.read.events")
	PermAppReadLog                       = PermissionRegistry.get("app.read.log")
	PermAppReadMetric                    = PermissionRegistry.get("app.read.metric")
	PermAppRun                           = PermissionRegistry.get("app.run")
	PermAppUpdate                        = PermissionRegistry.get("app.update")
	PermAppUpdateBind                    = PermissionRegistry.get("app.update.bind")
	PermAppUpdateCname                   = PermissionRegistry.get("app.update.cname")
	PermAppUpdateCnameAdd                = PermissionRegistry.get("app.update.cname.add")
	PermAppUpdateCnameRemove             = PermissionRegistry.get("app.update.cname.remove")
	PermAppUpdateDescription             = PermissionRegistry.get("app.update.description")
	PermAppUpdateEnv                     = PermissionRegistry.get("app.update.env")
	PermAppUpdateEnvSet                  = PermissionRegistry.get("app.update.env.set")
	PermAppUpdateEnvUnset                = PermissionRegistry.get("app.update.env.unset")
	PermAppUpdateGrant                   = PermissionRegistry.get("app.update.grant")
	PermAppUpdateLog                     = PermissionRegistry.get("app.update.log")
	PermAppUpdatePlan                    = PermissionRegistry.get("app.update.plan")
	PermAppUpdatePool                    = PermissionRegistry.get("app.update.pool")
	PermAppUpdateRestart                 = PermissionRegistry.get("app.update.restart")
	PermAppUpdateRevoke                  = PermissionRegistry.get("app.update.revoke")
	PermAppUpdateSleep                   = PermissionRegistry.get("app.update.sleep")
	PermAppUpdateStart                   = PermissionRegistry.get("app.update.start")
	PermAppUpdateStop                    = PermissionRegistry.get("app.update.stop")
	PermAppUpdateSwap                    = PermissionRegistry.get("app.update.swap")
	PermAppUpdateTeamowner               = PermissionRegistry.get("app.update.teamowner")
	PermAppUpdateUnbind                  = PermissionRegistry.get("app.update.unbind")
	PermAppUpdateUnit                    = PermissionRegistry.get("app.update.unit")
	PermAppUpdateUnitAdd                 = PermissionRegistry.get("app.update.unit.add")
	PermAppUpdateUnitRegister            = PermissionRegistry.get("app.update.unit.register")
	PermAppUpdateUnitRemove              = PermissionRegistry.get("app.update.unit.remove")
	PermAppUpdateUnitStatus              = PermissionRegistry.get("app.update.unit.status")
	PermDebug                            = PermissionRegistry.get("debug")
	PermHealing                          = PermissionRegistry.get("healing")
	PermHealingRead                      = PermissionRegistry.get("healing.read")
	PermHealingUpdate                    = PermissionRegistry.get("healing.update")
	PermMachine                          = PermissionRegistry.get("machine")
	PermMachineCreate                    = PermissionRegistry.get("machine.create")
	PermMachineDelete                    = PermissionRegistry.get("machine.delete")
	PermMachineRead                      = PermissionRegistry.get("machine.read")
	PermMachineTemplate                  = PermissionRegistry.get("machine.template")
	PermMachineTemplateCreate            = PermissionRegistry.get("machine.template.create")
	PermMachineTemplateDelete            = PermissionRegistry.get("machine.template.delete")
	PermMachineTemplateRead              = PermissionRegistry.get("machine.template.read")
	PermMachineTemplateUpdate            = PermissionRegistry.get("machine.template.update")
	PermNode                             = PermissionRegistry.get("node")
	PermNodeAutoscale                    = PermissionRegistry.get("node.autoscale")
	PermNodeCreate                       = PermissionRegistry.get("node.create")
	PermNodeDelete                       = PermissionRegistry.get("node.delete")
	PermNodeRead                         = PermissionRegistry.get("node.read")
	PermNodeReadEvents                   = PermissionRegistry.get("node.read.events")
	PermNodeUpdate                       = PermissionRegistry.get("node.update")
	PermNodecontainer                    = PermissionRegistry.get("nodecontainer")
	PermNodecontainerCreate              = PermissionRegistry.get("nodecontainer.create")
	PermNodecontainerDelete              = PermissionRegistry.get("nodecontainer.delete")
	PermNodecontainerRead                = PermissionRegistry.get("nodecontainer.read")
	PermNodecontainerUpdate              = PermissionRegistry.get("nodecontainer.update")
	PermNodecontainerUpdateUpgrade       = PermissionRegistry.get("nodecontainer.update.upgrade")
	PermPlan                             = PermissionRegistry.get("plan")
	PermPlanCreate                       = PermissionRegistry.get("plan.create")
	PermPlanDelete                       = PermissionRegistry.get("plan.delete")
	PermPlatform                         = PermissionRegistry.get("platform")
	PermPlatformCreate                   = PermissionRegistry.get("platform.create")
	PermPlatformDelete                   = PermissionRegistry.get("platform.delete")
	PermPlatformUpdate                   = PermissionRegistry.get("platform.update")
	PermPool                             = PermissionRegistry.get("pool")
	PermPoolCreate                       = PermissionRegistry.get("pool.create")
	PermPoolDelete                       = PermissionRegistry.get("pool.delete")
	PermPoolRead                         = PermissionRegistry.get("pool.read")
	PermPoolReadEvents                   = PermissionRegistry.get("pool.read.events")
	PermPoolUpdate                       = PermissionRegistry.get("pool.update")
	PermPoolUpdateLogs                   = PermissionRegistry.get("pool.update.logs")
	PermRole                             = PermissionRegistry.get("role")
	PermRoleCreate                       = PermissionRegistry.get("role.create")
	PermRoleDefault                      = PermissionRegistry.get("role.default")
	PermRoleDefaultCreate                = PermissionRegistry.get("role.default.create")
	PermRoleDefaultDelete                = PermissionRegistry.get("role.default.delete")
	PermRoleDelete                       = PermissionRegistry.get("role.delete")
	PermRoleUpdate                       = PermissionRegistry.get("role.update")
	PermRoleUpdateAssign                 = PermissionRegistry.get("role.update.assign")
	PermRoleUpdateDissociate             = PermissionRegistry.get("role.update.dissociate")
	PermService                          = PermissionRegistry.get("service")
	PermServiceInstance                  = PermissionRegistry.get("service-instance")
	PermServiceInstanceCreate            = PermissionRegistry.get("service-instance.create")
	PermServiceInstanceDelete            = PermissionRegistry.get("service-instance.delete")
	PermServiceInstanceRead              = PermissionRegistry.get("service-instance.read")
	PermServiceInstanceReadEvents        = PermissionRegistry.get("service-instance.read.events")
	PermServiceInstanceReadStatus        = PermissionRegistry.get("service-instance.read.status")
	PermServiceInstanceUpdate            = PermissionRegistry.get("service-instance.update")
	PermServiceInstanceUpdateBind        = PermissionRegistry.get("service-instance.update.bind")
	PermServiceInstanceUpdateDescription = PermissionRegistry.get("service-instance.update.description")
	PermServiceInstanceUpdateGrant       = PermissionRegistry.get("service-instance.update.grant")
	PermServiceInstanceUpdateProxy       = PermissionRegistry.get("service-instance.update.proxy")
	PermServiceInstanceUpdateRevoke      = PermissionRegistry.get("service-instance.update.revoke")
	PermServiceInstanceUpdateUnbind      = PermissionRegistry.get("service-instance.update.unbind")
	PermServiceCreate                    = PermissionRegistry.get("service.create")
	PermServiceDelete                    = PermissionRegistry.get("service.delete")
	PermServiceRead                      = PermissionRegistry.get("service.read")
	PermServiceReadDoc                   = PermissionRegistry.get("service.read.doc")
	PermServiceReadEvents                = PermissionRegistry.get("service.read.events")
	PermServiceReadPlans                 = PermissionRegistry.get("service.read.plans")
	PermServiceUpdate                    = PermissionRegistry.get("service.update")
	PermServiceUpdateDoc                 = PermissionRegistry.get("service.update.doc")
	PermServiceUpdateGrantAccess         = PermissionRegistry.get("service.update.grant-access")
	PermServiceUpdateProxy               = PermissionRegistry.get("service.update.proxy")
	PermServiceUpdateRevokeAccess        = PermissionRegistry.get("service.update.revoke-access")
	PermTeam                             = PermissionRegistry.get("team")
	PermTeamCreate                       = PermissionRegistry.get("team.create")
	PermTeamDelete                       = PermissionRegistry.get("team.delete")
	PermTeamRead                         = PermissionRegistry.get("team.read")
	PermTeamReadEvents                   = PermissionRegistry.get("team.read.events")
	PermUser                             = PermissionRegistry.get("user")
	PermUserCreate                       = PermissionRegistry.get("user.create")
	PermUserDelete                       = PermissionRegistry.get("user.delete")
	PermUserRead                         = PermissionRegistry.get("user.read")
	PermUserReadEvents                   = PermissionRegistry.get("user.read.events")
	PermUserUpdate                       = PermissionRegistry.get("user.update")
	PermUserUpdateKey                    = PermissionRegistry.get("user.update.key")
	PermUserUpdateKeyAdd                 = PermissionRegistry.get("user.update.key.add")
	PermUserUpdateKeyRemove              = PermissionRegistry.get("user.update.key.remove")
	PermUserUpdatePassword               = PermissionRegistry.get("user.update.password")
	PermUserUpdateQuota                  = PermissionRegistry.get("user.update.quota")
	PermUserUpdateReset                  = PermissionRegistry.get("user.update.reset")
	PermUserUpdateToken                  = PermissionRegistry.get("user.update.token")
)
