package mixeroauth

// ScopeCollection represents a slice of scopes
type ScopeCollection []Scope

// ScopeStringMap represents a map with string => scope pairs
type ScopeStringMap map[string]Scope

// Scope represents a scoped permission that can be listed when request OAUTH2 access to Mixer's API
type Scope string

const (

	// Let's you act as this user on other resources.
	UserActAs Scope = "user:act_as"

	// View your earned achievements.
	AchievementViewSelf Scope = "achievement:view:self"

	// View your channel analytics.
	ChannelAnalyticsSelf Scope = "channel:analytics:self"

	// Create new clips from videos on your channel.
	ChannelClipCreateSelf Scope = "channel:clip:create:self"

	// Allows deleting existing clips on your channel.
	ChannelClipDeleteSelf Scope = "channel:clip:delete:self"

	// Manage your costreaming requests.
	ChannelCostreamSelf Scope = "channel:costream:self"

	// Delete your channel banner
	ChannelDeleteBannerSelf Scope = "channel:deleteBanner:self"

	// View your channel details.
	ChannelDetailsSelf Scope = "channel:details:self"

	// Follow and unfollow other channels.
	ChannelFollowSelf Scope = "channel:follow:self"

	// Create and view partnership applications.
	ChannelPartnership Scope = "channel:partnership"

	// Manage your partnership status.
	ChannelPartnershipSelf Scope = "channel:partnership:self"

	// View your channel's stream key.
	ChannelStreamKeySelf Scope = "channel:streamKey:self"

	// Watch your test streams.
	ChannelTeststreamViewSelf Scope = "channel:teststream:view:self"

	// Update your channel settings
	ChannelUpdateSelf Scope = "channel:update:self"

	// Bypasses the catbot chat filter.
	ChatBypassCatbot Scope = "chat:bypass_catbot"

	// Bypass the chat content filter.
	ChatBypassFilter Scope = "chat:bypass_filter"

	// Bypass links being disallowed in chat.
	ChatBypassLinks Scope = "chat:bypass_links"

	// Bypass slowchat settings on channels.
	ChatBypassSlowchat Scope = "chat:bypass_slowchat"

	// Cancel a skill.
	ChatCancelSkill Scope = "chat:cancel_skill"

	// Manage bans in chats.
	ChatChangeBan Scope = "chat:change_ban"

	// Manage roles in chats.
	ChatChangeRole Scope = "chat:change_role"

	// Interact with chats on your behalf.
	ChatChat Scope = "chat:chat"

	// Clear messages in chats where authorized.
	ChatClearMessages Scope = "chat:clear_messages"

	// Connect to chat.
	ChatConnect Scope = "chat:connect"

	// Edit chat options, including links settings and slowchat.
	ChatEditOptions Scope = "chat:edit_options"

	// Start a giveaway in chats where authorized.
	ChatGiveawayStart Scope = "chat:giveaway_start"

	// Start a poll in chats where authorized.
	ChatPollStart Scope = "chat:poll_start"

	// Vote in chat polls.
	ChatPollVote Scope = "chat:poll_vote"

	// Clear all messages from a specific user in chat.
	ChatPurge Scope = "chat:purge"

	// Remove own and other's messages in chat.
	ChatRemoveMessage Scope = "chat:remove_message"

	// Change timeout settings in chats.
	ChatTimeout Scope = "chat:timeout"

	// View deleted messages in chat.
	ChatViewDeleted Scope = "chat:view_deleted"

	// Gives the ability to whisper in a channel
	ChatWhisper Scope = "chat:whisper"

	// View your Mixer homepage experience and recommendations.
	DelveViewSelf Scope = "delve:view:self"

	// Create, update and delete the interactive games in your account.
	InteractiveManageSelf Scope = "interactive:manage:self"

	// Run as an interactive game in your channel.
	InteractiveRobotSelf Scope = "interactive:robot:self"

	// View the users invoices.
	InvoiceViewSelf Scope = "invoice:view:self"

	// View and manage your security log.
	LogViewSelf Scope = "log:view:self"

	// View and manage your OAuth clients.
	OauthManageSelf Scope = "oauth:manage:self"

	// Manage the users VODs.
	RecordingManageSelf Scope = "recording:manage:self"

	// Create redeemables after performing a purchase.
	RedeemableCreateSelf Scope = "redeemable:create:self"

	// Use users redeemable.
	RedeemableRedeemSelf Scope = "redeemable:redeem:self"

	// View users redeemables.
	RedeemableViewSelf Scope = "redeemable:view:self"

	// View emoticons and other graphical resources you have access to.
	ResourceFindSelf Scope = "resource:find:self"

	// Cancel your subscriptions.
	SubscriptionCancelSelf Scope = "subscription:cancel:self"

	// Create new subscriptions.
	SubscriptionCreateSelf Scope = "subscription:create:self"

	// Renew your existing subscriptions.
	SubscriptionRenewSelf Scope = "subscription:renew:self"

	// View who you're subscribed to.
	SubscriptionViewSelf Scope = "subscription:view:self"

	// Administrate teams the user has rights in.
	TeamAdminister Scope = "team:administer"

	// Create, join, leave teams and set the users primary team.
	TeamManageSelf Scope = "team:manage:self"

	// Cancel pending transactions.
	TransactionCancelSelf Scope = "transaction:cancel:self"

	// View your pending transactions.
	TransactionViewSelf Scope = "transaction:view:self"

	// View your user analytics
	UserAnalyticsSelf Scope = "user:analytics:self"

	// View your email address and other private details.
	UserDetailsSelf Scope = "user:details:self"

	// View users discord invites.
	UserGetDiscordInviteSelf Scope = "user:getDiscordInvite:self"

	// View your user security log.
	UserLogSelf Scope = "user:log:self"

	// View and manage your notifications.
	UserNotificationSelf Scope = "user:notification:self"

	// Mark a VOD as seen for the user.
	UserSeenSelf Scope = "user:seen:self"

	// Update your account, including your email but not your password.
	UserUpdateSelf Scope = "user:update:self"

	// Update your password.
	UserUpdatePasswordSelf Scope = "user:updatePassword:self"
)
