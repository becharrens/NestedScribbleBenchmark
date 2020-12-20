package protocol

import "NestedScribbleBenchmark/quicksort/messages"
import "NestedScribbleBenchmark/quicksort/channels/quicksort2"
import "NestedScribbleBenchmark/quicksort/channels/quicksort"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import quicksort_2 "NestedScribbleBenchmark/quicksort/results/quicksort"
import "NestedScribbleBenchmark/quicksort/roles"
import "sync"

type QuickSort_Env interface {
	New_Partition_Env() callbacks.QuickSort_Partition_Env
	New_Left_Env() callbacks.QuickSort_Left_Env
	New_Right_Env() callbacks.QuickSort_Right_Env
	Partition_Result(result quicksort_2.Partition_Result)
	Left_Result(result quicksort_2.Left_Result)
	Right_Result(result quicksort_2.Right_Result)
}

func Start_QuickSort_Partition(protocolEnv QuickSort_Env, wg *sync.WaitGroup, roleChannels quicksort.Partition_Chan, inviteChannels invitations.QuickSort_Partition_InviteChan, env callbacks.QuickSort_Partition_Env) {
	defer wg.Done()
	result := roles.QuickSort_Partition(wg, roleChannels, inviteChannels, env)
	protocolEnv.Partition_Result(result)
}

func Start_QuickSort_Left(protocolEnv QuickSort_Env, wg *sync.WaitGroup, roleChannels quicksort.Left_Chan, inviteChannels invitations.QuickSort_Left_InviteChan, env callbacks.QuickSort_Left_Env) {
	defer wg.Done()
	result := roles.QuickSort_Left(wg, roleChannels, inviteChannels, env)
	protocolEnv.Left_Result(result)
}

func Start_QuickSort_Right(protocolEnv QuickSort_Env, wg *sync.WaitGroup, roleChannels quicksort.Right_Chan, inviteChannels invitations.QuickSort_Right_InviteChan, env callbacks.QuickSort_Right_Env) {
	defer wg.Done()
	result := roles.QuickSort_Right(wg, roleChannels, inviteChannels, env)
	protocolEnv.Right_Result(result)
}

func QuickSort(protocolEnv QuickSort_Env) {
	right_partition_intarr := make(chan []int, 1)
	right_partition_label := make(chan messages.QuickSort_Label, 1)
	left_partition_intarr := make(chan []int, 1)
	left_partition_label := make(chan messages.QuickSort_Label, 1)
	right_invite_right := make(chan quicksort2.P_Chan, 1)
	right_invite_right_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	left_invite_left := make(chan quicksort2.P_Chan, 1)
	left_invite_left_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	partition_right_intarr := make(chan []int, 1)
	partition_right_label := make(chan messages.QuickSort_Label, 1)
	partition_left_intarr := make(chan []int, 1)
	partition_left_label := make(chan messages.QuickSort_Label, 1)

	right_chan := quicksort.Right_Chan{
		Label_To_Partition:    right_partition_label,
		Label_From_Partition:  partition_right_label,
		IntArr_To_Partition:   right_partition_intarr,
		IntArr_From_Partition: partition_right_intarr,
	}
	partition_chan := quicksort.Partition_Chan{
		Label_To_Right:    partition_right_label,
		Label_To_Left:     partition_left_label,
		Label_From_Right:  right_partition_label,
		Label_From_Left:   left_partition_label,
		IntArr_To_Right:   partition_right_intarr,
		IntArr_To_Left:    partition_left_intarr,
		IntArr_From_Right: right_partition_intarr,
		IntArr_From_Left:  left_partition_intarr,
	}
	left_chan := quicksort.Left_Chan{
		Label_To_Partition:    left_partition_label,
		Label_From_Partition:  partition_left_label,
		IntArr_To_Partition:   left_partition_intarr,
		IntArr_From_Partition: partition_left_intarr,
	}

	right_inviteChan := invitations.QuickSort_Right_InviteChan{
		Invite_Right_To_QuickSort2_P_InviteChan: right_invite_right_invitechan,
		Invite_Right_To_QuickSort2_P:            right_invite_right,
	}
	partition_inviteChan := invitations.QuickSort_Partition_InviteChan{}
	left_inviteChan := invitations.QuickSort_Left_InviteChan{
		Invite_Left_To_QuickSort2_P_InviteChan: left_invite_left_invitechan,
		Invite_Left_To_QuickSort2_P:            left_invite_left,
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
