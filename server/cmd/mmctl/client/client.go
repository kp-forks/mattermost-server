// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package client

import (
	"context"
	"io"
	"net/http"

	"github.com/mattermost/mattermost/server/public/model"
)

type Client interface {
	CreateChannel(ctx context.Context, channel *model.Channel) (*model.Channel, *model.Response, error)
	RemoveUserFromChannel(ctx context.Context, channelID, userID string) (*model.Response, error)
	GetChannelMembers(ctx context.Context, channelID string, page, perPage int, etag string) (model.ChannelMembers, *model.Response, error)
	AddChannelMember(ctx context.Context, channelID, userID string) (*model.ChannelMember, *model.Response, error)
	DeleteChannel(ctx context.Context, channelID string) (*model.Response, error)
	PermanentDeleteChannel(ctx context.Context, channelID string) (*model.Response, error)
	MoveChannel(ctx context.Context, channelID, teamID string, force bool) (*model.Channel, *model.Response, error)
	GetPublicChannelsForTeam(ctx context.Context, teamID string, page int, perPage int, etag string) ([]*model.Channel, *model.Response, error)
	GetDeletedChannelsForTeam(ctx context.Context, teamID string, page int, perPage int, etag string) ([]*model.Channel, *model.Response, error)
	GetPrivateChannelsForTeam(ctx context.Context, teamID string, page int, perPage int, etag string) ([]*model.Channel, *model.Response, error)
	GetChannelsForTeamForUser(ctx context.Context, teamID, userID string, includeDeleted bool, etag string) ([]*model.Channel, *model.Response, error)
	RestoreChannel(ctx context.Context, channelID string) (*model.Channel, *model.Response, error)
	PatchChannel(ctx context.Context, channelID string, patch *model.ChannelPatch) (*model.Channel, *model.Response, error)
	GetChannelByName(ctx context.Context, channelName, teamID string, etag string) (*model.Channel, *model.Response, error)
	GetChannelByNameIncludeDeleted(ctx context.Context, channelName, teamID string, etag string) (*model.Channel, *model.Response, error)
	GetChannel(ctx context.Context, channelID, etag string) (*model.Channel, *model.Response, error)
	GetTeam(ctx context.Context, teamID, etag string) (*model.Team, *model.Response, error)
	GetTeamByName(ctx context.Context, name, etag string) (*model.Team, *model.Response, error)
	GetAllTeams(ctx context.Context, etag string, page int, perPage int) ([]*model.Team, *model.Response, error)
	CreateTeam(ctx context.Context, team *model.Team) (*model.Team, *model.Response, error)
	PatchTeam(ctx context.Context, teamID string, patch *model.TeamPatch) (*model.Team, *model.Response, error)
	AddTeamMember(ctx context.Context, teamID, userID string) (*model.TeamMember, *model.Response, error)
	RemoveTeamMember(ctx context.Context, teamID, userID string) (*model.Response, error)
	SoftDeleteTeam(ctx context.Context, teamID string) (*model.Response, error)
	PermanentDeleteTeam(ctx context.Context, teamID string) (*model.Response, error)
	RestoreTeam(ctx context.Context, teamID string) (*model.Team, *model.Response, error)
	UpdateTeamPrivacy(ctx context.Context, teamID string, privacy string) (*model.Team, *model.Response, error)
	SearchTeams(ctx context.Context, search *model.TeamSearch) ([]*model.Team, *model.Response, error)
	GetPost(ctx context.Context, postID string, etag string) (*model.Post, *model.Response, error)
	CreatePost(ctx context.Context, post *model.Post) (*model.Post, *model.Response, error)
	GetPostsForChannel(ctx context.Context, channelID string, page, perPage int, etag string, collapsedThreads bool, includeDeleted bool) (*model.PostList, *model.Response, error)
	GetPostsSince(ctx context.Context, channelID string, since int64, collapsedThreads bool) (*model.PostList, *model.Response, error)
	DoAPIPost(ctx context.Context, url string, data string) (*http.Response, error)
	GetLdapGroups(ctx context.Context) ([]*model.Group, *model.Response, error)
	GetGroupsByChannel(ctx context.Context, channelID string, groupOpts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, int, *model.Response, error)
	GetGroupsByTeam(ctx context.Context, teamID string, groupOpts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, int, *model.Response, error)
	RestoreGroup(ctx context.Context, groupID string, etag string) (*model.Group, *model.Response, error)
	UploadLicenseFile(ctx context.Context, data []byte) (*model.Response, error)
	RemoveLicenseFile(ctx context.Context) (*model.Response, error)
	GetLogs(ctx context.Context, page, perPage int) ([]string, *model.Response, error)
	GetRoleByName(ctx context.Context, name string) (*model.Role, *model.Response, error)
	PatchRole(ctx context.Context, roleID string, patch *model.RolePatch) (*model.Role, *model.Response, error)
	UploadPlugin(ctx context.Context, file io.Reader) (*model.Manifest, *model.Response, error)
	UploadPluginForced(ctx context.Context, file io.Reader) (*model.Manifest, *model.Response, error)
	RemovePlugin(ctx context.Context, id string) (*model.Response, error)
	EnablePlugin(ctx context.Context, id string) (*model.Response, error)
	DisablePlugin(ctx context.Context, id string) (*model.Response, error)
	GetPlugins(ctx context.Context) (*model.PluginsResponse, *model.Response, error)
	GetUser(ctx context.Context, userID, etag string) (*model.User, *model.Response, error)
	GetUserByUsername(ctx context.Context, userName, etag string) (*model.User, *model.Response, error)
	GetUserByEmail(ctx context.Context, email, etag string) (*model.User, *model.Response, error)
	GetUsersByIds(ctx context.Context, userIDs []string) ([]*model.User, *model.Response, error)
	GetUsersWithCustomQueryParameters(ctx context.Context, page, perPage int, queryParameters string, etag string) ([]*model.User, *model.Response, error)
	GetUsersInTeam(ctx context.Context, teamID string, page, perPage int, etag string) ([]*model.User, *model.Response, error)
	PermanentDeleteUser(ctx context.Context, userID string) (*model.Response, error)
	PermanentDeleteAllUsers(ctx context.Context) (*model.Response, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, *model.Response, error)
	VerifyUserEmailWithoutToken(ctx context.Context, userID string) (*model.User, *model.Response, error)
	UpdateUserRoles(ctx context.Context, userID, roles string) (*model.Response, error)
	InviteUsersToTeam(ctx context.Context, teamID string, userEmails []string) (*model.Response, error)
	SendPasswordResetEmail(ctx context.Context, email string) (*model.Response, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, *model.Response, error)
	UpdateUserAuth(ctx context.Context, userId string, userAuth *model.UserAuth) (*model.UserAuth, *model.Response, error)
	UpdateUserMfa(ctx context.Context, userID, code string, activate bool) (*model.Response, error)
	UpdateUserPassword(ctx context.Context, userID, currentPassword, newPassword string) (*model.Response, error)
	UpdateUserHashedPassword(ctx context.Context, userID, newHashedPassword string) (*model.Response, error)
	CreateUserAccessToken(ctx context.Context, userID, description string) (*model.UserAccessToken, *model.Response, error)
	RevokeUserAccessToken(ctx context.Context, tokenID string) (*model.Response, error)
	GetUserAccessTokensForUser(ctx context.Context, userID string, page, perPage int) ([]*model.UserAccessToken, *model.Response, error)
	ConvertUserToBot(ctx context.Context, userID string) (*model.Bot, *model.Response, error)
	ConvertBotToUser(ctx context.Context, userID string, userPatch *model.UserPatch, setSystemAdmin bool) (*model.User, *model.Response, error)
	PromoteGuestToUser(ctx context.Context, userID string) (*model.Response, error)
	DemoteUserToGuest(ctx context.Context, guestID string) (*model.Response, error)
	CreateCommand(ctx context.Context, cmd *model.Command) (*model.Command, *model.Response, error)
	ListCommands(ctx context.Context, teamID string, customOnly bool) ([]*model.Command, *model.Response, error)
	GetCommandById(ctx context.Context, cmdID string) (*model.Command, *model.Response, error)
	UpdateCommand(ctx context.Context, cmd *model.Command) (*model.Command, *model.Response, error)
	MoveCommand(ctx context.Context, teamID string, commandID string) (*model.Response, error)
	DeleteCommand(ctx context.Context, commandID string) (*model.Response, error)
	GetConfig(ctx context.Context) (*model.Config, *model.Response, error)
	GetConfigWithOptions(ctx context.Context, options model.GetConfigOptions) (map[string]any, *model.Response, error)
	GetOldClientConfig(ctx context.Context, etag string) (map[string]string, *model.Response, error)
	UpdateConfig(context.Context, *model.Config) (*model.Config, *model.Response, error)
	PatchConfig(context.Context, *model.Config) (*model.Config, *model.Response, error)
	ReloadConfig(ctx context.Context) (*model.Response, error)
	MigrateConfig(ctx context.Context, from, to string) (*model.Response, error)
	SyncLdap(ctx context.Context) (*model.Response, error)
	MigrateIdLdap(ctx context.Context, toAttribute string) (*model.Response, error)
	GetUsers(ctx context.Context, page, perPage int, etag string) ([]*model.User, *model.Response, error)
	UpdateUserActive(ctx context.Context, userID string, activate bool) (*model.Response, error)
	UpdateTeam(ctx context.Context, team *model.Team) (*model.Team, *model.Response, error)
	UpdateChannelPrivacy(ctx context.Context, channelID string, privacy model.ChannelType) (*model.Channel, *model.Response, error)
	CreateBot(ctx context.Context, bot *model.Bot) (*model.Bot, *model.Response, error)
	PatchBot(ctx context.Context, userID string, patch *model.BotPatch) (*model.Bot, *model.Response, error)
	GetBots(ctx context.Context, page, perPage int, etag string) ([]*model.Bot, *model.Response, error)
	GetBotsIncludeDeleted(ctx context.Context, page, perPage int, etag string) ([]*model.Bot, *model.Response, error)
	GetBotsOrphaned(ctx context.Context, page, perPage int, etag string) ([]*model.Bot, *model.Response, error)
	DisableBot(ctx context.Context, botUserID string) (*model.Bot, *model.Response, error)
	EnableBot(ctx context.Context, botUserID string) (*model.Bot, *model.Response, error)
	AssignBot(ctx context.Context, botUserID, newOwnerID string) (*model.Bot, *model.Response, error)
	SetServerBusy(ctx context.Context, secs int) (*model.Response, error)
	ClearServerBusy(ctx context.Context) (*model.Response, error)
	GetServerBusy(ctx context.Context) (*model.ServerBusyState, *model.Response, error)
	CheckIntegrity(ctx context.Context) ([]model.IntegrityCheckResult, *model.Response, error)
	InstallPluginFromURL(context.Context, string, bool) (*model.Manifest, *model.Response, error)
	InstallMarketplacePlugin(context.Context, *model.InstallMarketplacePluginRequest) (*model.Manifest, *model.Response, error)
	GetMarketplacePlugins(context.Context, *model.MarketplacePluginFilter) ([]*model.MarketplacePlugin, *model.Response, error)
	MigrateAuthToLdap(ctx context.Context, fromAuthService string, matchField string, force bool) (*model.Response, error)
	MigrateAuthToSaml(ctx context.Context, fromAuthService string, usersMap map[string]string, auto bool) (*model.Response, error)
	GetPing(ctx context.Context) (string, *model.Response, error)
	GetPingWithFullServerStatus(ctx context.Context) (map[string]string, *model.Response, error)
	GetPingWithOptions(ctx context.Context, options model.SystemPingOptions) (map[string]string, *model.Response, error)
	CreateUpload(ctx context.Context, us *model.UploadSession) (*model.UploadSession, *model.Response, error)
	GetUpload(ctx context.Context, uploadID string) (*model.UploadSession, *model.Response, error)
	GetUploadsForUser(ctx context.Context, userID string) ([]*model.UploadSession, *model.Response, error)
	UploadData(ctx context.Context, uploadID string, data io.Reader) (*model.FileInfo, *model.Response, error)
	ListImports(ctx context.Context) ([]string, *model.Response, error)
	DeleteImport(ctx context.Context, name string) (*model.Response, error)
	GetJob(ctx context.Context, id string) (*model.Job, *model.Response, error)
	GetJobs(ctx context.Context, jobType string, status string, page int, perPage int) ([]*model.Job, *model.Response, error)
	GetJobsByType(ctx context.Context, jobType string, page int, perPage int) ([]*model.Job, *model.Response, error)
	CreateJob(ctx context.Context, job *model.Job) (*model.Job, *model.Response, error)
	CancelJob(ctx context.Context, jobID string) (*model.Response, error)
	UpdateJobStatus(ctx context.Context, jobId string, status string, force bool) (*model.Response, error)
	CreateIncomingWebhook(ctx context.Context, hook *model.IncomingWebhook) (*model.IncomingWebhook, *model.Response, error)
	UpdateIncomingWebhook(ctx context.Context, hook *model.IncomingWebhook) (*model.IncomingWebhook, *model.Response, error)
	GetIncomingWebhooks(ctx context.Context, page int, perPage int, etag string) ([]*model.IncomingWebhook, *model.Response, error)
	GetIncomingWebhooksForTeam(ctx context.Context, teamID string, page int, perPage int, etag string) ([]*model.IncomingWebhook, *model.Response, error)
	GetIncomingWebhook(ctx context.Context, hookID string, etag string) (*model.IncomingWebhook, *model.Response, error)
	DeleteIncomingWebhook(ctx context.Context, hookID string) (*model.Response, error)
	CreateOutgoingWebhook(ctx context.Context, hook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.Response, error)
	UpdateOutgoingWebhook(ctx context.Context, hook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhooks(ctx context.Context, page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhook(ctx context.Context, hookID string) (*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhooksForChannel(ctx context.Context, channelID string, page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhooksForTeam(ctx context.Context, teamID string, page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response, error)
	RegenOutgoingHookToken(ctx context.Context, hookID string) (*model.OutgoingWebhook, *model.Response, error)
	DeleteOutgoingWebhook(ctx context.Context, hookID string) (*model.Response, error)
	ListExports(ctx context.Context) ([]string, *model.Response, error)
	DeleteExport(ctx context.Context, name string) (*model.Response, error)
	DownloadExport(ctx context.Context, name string, wr io.Writer, offset int64) (int64, *model.Response, error)
	DownloadComplianceExport(ctx context.Context, jobID string, wr io.Writer) (string, error)
	GeneratePresignedURL(ctx context.Context, name string) (*model.PresignURLResponse, *model.Response, error)
	ResetSamlAuthDataToEmail(ctx context.Context, includeDeleted bool, dryRun bool, userIDs []string) (int64, *model.Response, error)
	GenerateSupportPacket(ctx context.Context) (io.ReadCloser, string, *model.Response, error)
	GetOAuthApps(ctx context.Context, page, perPage int) ([]*model.OAuthApp, *model.Response, error)
	GetPreferences(ctx context.Context, userId string) (model.Preferences, *model.Response, error)
	GetPreferencesByCategory(ctx context.Context, userId, category string) (model.Preferences, *model.Response, error)
	GetPreferenceByCategoryAndName(ctx context.Context, userId, category, preferenceName string) (*model.Preference, *model.Response, error)
	UpdatePreferences(ctx context.Context, userId string, preferences model.Preferences) (*model.Response, error)
	DeletePreferences(ctx context.Context, userId string, preferences model.Preferences) (*model.Response, error)
	PermanentDeletePost(ctx context.Context, postID string) (*model.Response, error)
	DeletePost(ctx context.Context, postId string) (*model.Response, error)
}
