package protocol

import "ScribbleBenchmark/quicksort/messages/quicksort"
import "ScribbleBenchmark/quicksort/channels/quicksort2"
import quicksort_2 "ScribbleBenchmark/quicksort/channels/quicksort"
import "ScribbleBenchmark/quicksort/invitations"
import "ScribbleBenchmark/quicksort/callbacks"
import quicksort_3 "ScribbleBenchmark/quicksort/results/quicksort"
import "ScribbleBenchmark/quicksort/roles"
import "sync"

type QuickSort_Env interface {
	New_Partition_Env() callbacks.QuickSort_Partition_Env
	New_Left_Env() callbacks.QuickSort_Left_Env
	New_Right_Env() callbacks.QuickSort_Right_Env
	Partition_Result(result quicksort_3.Partition_Result) 
	Left_Result(result quicksort_3.Left_Result) 
	Right_Result(result quicksort_3.Right_Result) 
}

func Start_QuickSort_Partition(protocolEnv QuickSort_Env, wg *sync.WaitGroup, roleChannels quicksort_2.Partition_Chan, inviteChannels invitations.QuickSort_Partition_InviteChan, env callbacks.QuickSort_Partition_Env)  {
	defer wg.Done()
	result := roles.QuickSort_Partition(wg, roleChannels, inviteChannels, env)
	protocolEnv.Partition_Result(result)
} 

func Start_QuickSort_Left(protocolEnv QuickSort_Env, wg *sync.WaitGroup, roleChannels quicksort_2.Left_Chan, inviteChannels invitations.QuickSort_Left_InviteChan, env callbacks.QuickSort_Left_Env)  {
	defer wg.Done()
	result := roles.QuickSort_Left(wg, roleChannels, inviteChannels, env)
	protocolEnv.Left_Result(result)
} 

func Start_QuickSort_Right(protocolEnv QuickSort_Env, wg *sync.WaitGroup, roleChannels quicksort_2.Right_Chan, inviteChannels invitations.QuickSort_Right_InviteChan, env callbacks.QuickSort_Right_Env)  {
	defer wg.Done()
	result := roles.QuickSort_Right(wg, roleChannels, inviteChannels, env)
	protocolEnv.Right_Result(result)
} 

func QuickSort(protocolEnv QuickSort_Env)  {
	partition_right_done := make(chan quicksort.Done, 1)
	partition_left_done := make(chan quicksort.Done, 1)
	right_partition_sortedright := make(chan quicksort.SortedRight, 1)
	left_partition_sortedleft := make(chan quicksort.SortedLeft, 1)
	right_invite_right := make(chan quicksort2.P_Chan, 1)
	right_invite_right_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	left_invite_left := make(chan quicksort2.P_Chan, 1)
	left_invite_left_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	partition_right_rightpartition := make(chan quicksort.RightPartition, 1)
	partition_left_leftparitition := make(chan quicksort.LeftParitition, 1)

	right_chan := quicksort_2.Right_Chan{
		Partition_SortedRight: right_partition_sortedright,
		Partition_RightPartition: partition_right_rightpartition,
		Partition_Done: partition_right_done,
	}
	partition_chan := quicksort_2.Partition_Chan{
		Right_SortedRight: right_partition_sortedright,
		Right_RightPartition: partition_right_rightpartition,
		Right_Done: partition_right_done,
		Left_SortedLeft: left_partition_sortedleft,
		Left_LeftParitition: partition_left_leftparitition,
		Left_Done: partition_left_done,
	}
	left_chan := quicksort_2.Left_Chan{
		Partition_SortedLeft: left_partition_sortedleft,
		Partition_LeftParitition: partition_left_leftparitition,
		Partition_Done: partition_left_done,
	}

	right_inviteChan := invitations.QuickSort_Right_InviteChan{
		Invite_Right_To_QuickSort2_P_InviteChan: right_invite_right_invitechan,
		Invite_Right_To_QuickSort2_P: right_invite_right,
	}
	partition_inviteChan := invitations.QuickSort_Partition_InviteChan{

	}
	left_inviteChan := invitations.QuickSort_Left_InviteChan{
		Invite_Left_To_QuickSort2_P_InviteChan: left_invite_left_invitechan,
		Invite_Left_To_QuickSort2_P: left_invite_left,
	}

	var wg sync.WaitGroup

	wg.Add(3)

	partition_env := protocolEnv.New_Partition_Env()
	left_env := protocolEnv.New_Left_Env()
	right_env := protocolEnv.New_Right_Env()

	go Start_QuickSort_Partition(protocolEnv, &wg, partition_chan, partition_inviteChan, partition_env)
	go Start_QuickSort_Left(protocolEnv, &wg, left_chan, left_inviteChan, left_env)
	go Start_QuickSort_Right(protocolEnv, &wg, right_chan, right_inviteChan, right_env)

	wg.Wait()
} 