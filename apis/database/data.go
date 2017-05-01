package database

import (
	// Stdlib
	"encoding/json"
	"strconv"
	"strings"

	// RPC
	"github.com/shaunmza/steemgo/types"
)

type Config struct {
	SteemitBlockchainHardforkVersion string `json:"STEEMIT_BLOCKCHAIN_HARDFORK_VERSION"`
	SteemitBlockchainVersion         string `json:"STEEMIT_BLOCKCHAIN_VERSION"`
	SteemitBlockInterval             uint   `json:"STEEMIT_BLOCK_INTERVAL"`
}

type DynamicGlobalProperties struct {
	Time                     *types.Time  `json:"time"`
	TotalPow                 *types.Int   `json:"total_pow"`
	NumPowWitnesses          *types.Int   `json:"num_pow_witnesses"`
	CurrentReserveRatio      *types.Int   `json:"current_reserve_ratio"`
	ID                       *types.ID    `json:"id"`
	CurrentSupply            string       `json:"current_supply"`
	CurrentSBDSupply         string       `json:"current_sbd_supply"`
	MaximumBlockSize         *types.Int   `json:"maximum_block_size"`
	RecentSlotsFilled        *types.Int   `json:"recent_slots_filled"`
	CurrentWitness           string       `json:"current_witness"`
	TotalRewardShares2       *types.Int   `json:"total_reward_shares2"`
	AverageBlockSize         *types.Int   `json:"average_block_size"`
	CurrentAslot             *types.Int   `json:"current_aslot"`
	LastIrreversibleBlockNum uint32       `json:"last_irreversible_block_num"`
	TotalVestingShares       string       `json:"total_vesting_shares"`
	TotalVersingFundSteem    string       `json:"total_vesting_fund_steem"`
	HeadBlockID              string       `json:"head_block_id"`
	HeadBlockNumber          types.UInt32 `json:"head_block_number"`
	VirtualSupply            string       `json:"virtual_supply"`
	ConfidentialSupply       string       `json:"confidential_supply"`
	ConfidentialSBDSupply    string       `json:"confidential_sbd_supply"`
	TotalRewardFundSteem     string       `json:"total_reward_fund_steem"`
	TotalActivityFundSteem   string       `json:"total_activity_fund_steem"`
	TotalActivityFundShares  *types.Int   `json:"total_activity_fund_shares"`
	SBDInterestRate          *types.Int   `json:"sbd_interest_rate"`
	MaxVirtualBandwidth      *types.Int   `json:"max_virtual_bandwidth"`
}

type Block struct {
	Number                uint32               `json:"-"`
	Timestamp             *types.Time          `json:"timestamp"`
	Witness               string               `json:"witness"`
	WitnessSignature      string               `json:"witness_signature"`
	TransactionMerkleRoot string               `json:"transaction_merkle_root"`
	Previous              string               `json:"previous"`
	Extensions            [][]interface{}      `json:"extensions"`
	Transactions          []*types.Transaction `json:"transactions"`
}

type Content struct {
	Id                      *types.ID        `json:"id"`
	RootTitle               string           `json:"root_title"`
	Active                  *types.Time      `json:"active"`
	AbsRshares              *types.Int       `json:"abs_rshares"`
	PendingPayoutValue      string           `json:"pending_payout_value"`
	TotalPendingPayoutValue string           `json:"total_pending_payout_value"`
	Category                string           `json:"category"`
	Title                   string           `json:"title"`
	LastUpdate              *types.Time      `json:"last_update"`
	Stats                   string           `json:"stats"`
	Body                    string           `json:"body"`
	Created                 *types.Time      `json:"created"`
	Replies                 []*Content       `json:"replies"`
	Permlink                string           `json:"permlink"`
	JsonMetadata            *ContentMetadata `json:"json_metadata"`
	Children                *types.Int       `json:"children"`
	NetRshares              *types.Int       `json:"net_rshares"`
	URL                     string           `json:"url"`
	ActiveVotes             []*VoteState     `json:"active_votes"`
	ParentPermlink          string           `json:"parent_permlink"`
	CashoutTime             *types.Time      `json:"cashout_time"`
	TotalPayoutValue        string           `json:"total_payout_value"`
	ParentAuthor            string           `json:"parent_author"`
	ChildrenRshares2        *types.Int       `json:"children_rshares2"`
	Author                  string           `json:"author"`
	Depth                   *types.Int       `json:"depth"`
	TotalVoteWeight         *types.Int       `json:"total_vote_weight"`
}

func (content *Content) IsStory() bool {
	return content.ParentAuthor == ""
}

type ContentMetadata struct {
	Flag  bool
	Users []string
	Tags  []string
	Image []string
}

type ContentMetadataRaw struct {
	Users types.StringSlice `json:"users"`
	Tags  types.StringSlice `json:"tags"`
	Image types.StringSlice `json:"image"`
}

func (metadata *ContentMetadata) UnmarshalJSON(data []byte) error {
	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	switch unquoted {
	case "true":
		metadata.Flag = true
		return nil
	case "false":
		metadata.Flag = false
		return nil
	}

	if len(unquoted) == 0 {
		var value ContentMetadata
		metadata = &value
		return nil
	}

	var raw ContentMetadataRaw
	if err := json.NewDecoder(strings.NewReader(unquoted)).Decode(&raw); err != nil {
		return err
	}

	metadata.Users = raw.Users
	metadata.Tags = raw.Tags
	metadata.Image = raw.Image

	return nil
}

type VoteState struct {
	Voter   string      `json:"voter"`
	Weight  *types.Int  `json:"weight"`
	Rshares *types.Int  `json:"rshares"`
	Percent *types.Int  `json:"percent"`
	Time    *types.Time `json:"time"`
}

type Account struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
	//Owner ??? 	`json:"owner"`
	//`json:"active":{},
	//`json:"posting":{},
	MemoKey string `json:"memo_key"`
	//json_metadata ??? `json:"json_metadata":"{\"profile\":{\"name\":\"Shaun\",\"about\":\"Developer, tinkerer, short attention span\",\"location\":\"Portugal\",\"profile_image\":\"http://i.imgur.com/wftoPVi.jpg\"}}",
	Proxy               string      `json:"proxy"`
	LastOwnerUpdate     *types.Time `json:"last_owner_update"`
	LastAccountUpdate   *types.Time `json:"last_account_update"`
	Created             *types.Time `json:"created"`
	Mined               bool        `json:"mined"`
	OwnerChallenged     bool        `json:"owner_challenged"`
	ActiveChallenged    bool        `json:"active_challenged"`
	LastOwnerProved     *types.Time `json:"last_owner_proved"`
	LastActiveProved    *types.Time `json:"last_active_proved"`
	RecoveryAccount     string      `json:"recovery_account"`
	LastAccountRecovery *types.Time `json:"last_account_recovery"`
	ResetAccount        string      `json:"reset_account"`
	CommentCount        uint32      `json:"comment_count"`
	LifetimeVoteCount   uint32      `json:"lifetime_vote_count"`
	PostCount           uint32      `json:"post_count"`
	CanVote             bool        `json:"can_vote"`
	VotingPower         uint32      `json:"voting_power"`
	LastVoteTime        *types.Time `json:"last_vote_time"`
	Balance             string      `json:"balance"`
	SavingsBalance      string      `json:"savings_balance"`
	SbdBalance          string      `json:"sbd_balance"`
	//`json:"sbd_seconds":"227233098",
	SbdSecondsLastUpdate          *types.Time `json:"sbd_seconds_last_update"`
	SbdLastInterestPayment        *types.Time `json:"sbd_last_interest_payment"`
	SavingsSbdBalance             string      `json:"savings_sbd_balance"`
	SavingsSbdSeconds             string      `json:"savings_sbd_seconds"`
	SavingsSbdSecondsLastUpdate   *types.Time `json:"savings_sbd_seconds_last_update"`
	SavingsSbdLastInterestPayment *types.Time `json:"savings_sbd_last_interest_payment"`
	SavingsWithdrawRequests       uint32      `json:"savings_withdraw_requests"`
	RewardSbdBalance              string      `json:"reward_sbd_balance"`
	RewardSteemBalance            string      `json:"reward_steem_balance"`
	RewardVestingBalance          string      `json:"reward_vesting_balance"`
	RewardVestingSteem            string      `json:"reward_vesting_steem"`
	VestingShares                 string      `json:"vesting_shares"`
	DelegatedVestingShares        string      `json:"delegated_vesting_shares"`
	ReceivedVestingShares         string      `json:"received_vesting_shares"`
	VestingWithdrawRate           string      `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal         *types.Time `json:"next_vesting_withdrawal"`
	Withdrawn                     uint32      `json:"withdrawn"`
	ToWithdraw                    uint32      `json:"to_withdraw"`
	WithdrawRoutes                uint32      `json:"withdraw_routes"`
	CurationRewards               uint32      `json:"curation_rewards"`
	PostingRewards                uint32      `json:"posting_rewards"`
	//proxied_vsf_votes `json:"proxied_vsf_votes":[],
	WitnessesVotedFor         uint32      `json:"witnesses_voted_for"`
	AverageBandwidth          uint32      `json:"average_bandwidth"`
	LifetimeBandwidth         string      `json:"lifetime_bandwidth"`
	LastBandwidthUpdate       *types.Time `json:"last_bandwidth_update"`
	AverageMarketBandwidth    uint32      `json:"average_market_bandwidth"`
	LastMarketBandwidthUpdate *types.Time `json:"last_market_bandwidth_update"`
	LastPost                  *types.Time `json:"last_post"`
	LastRootPost              *types.Time `json:"last_root_post"`
	PostBandwidth             uint32      `json:"post_bandwidth"`
	NewAverageBandwidth       uint32      `json:"new_average_bandwidth,uint32"`
	NewAverageMarketBandwidth string      `json:"new_average_market_bandwidth"`
	VestingBalance            string      `json:"vesting_balance"`
	Reputation                string      `json:"reputation"`
	//`json:"transfer_history":[],
	//`json:"market_history":[],
	//`json:"post_history":[],
	//`json:"vote_history":[],
	//`json:"other_history":[],
	//`json:"witness_votes":[],
	//`json:"tags_usage":[],
	//`json:"guest_bloggers":[],
	//`json:"blog_category":{}
}
