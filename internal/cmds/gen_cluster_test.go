// Code generated DO NOT EDIT

package cmds

import "testing"

func cluster0(s Builder) {
	s.Asking().Build()
	s.ClusterAddslots().Slot(1).Slot(1).Build()
	s.ClusterAddslotsrange().StartSlotEndSlot().StartSlotEndSlot(1, 1).StartSlotEndSlot(1, 1).Build()
	s.ClusterBumpepoch().Build()
	s.ClusterCountFailureReports().NodeId("1").Build()
	s.ClusterCountkeysinslot().Slot(1).Build()
	s.ClusterDelslots().Slot(1).Slot(1).Build()
	s.ClusterDelslotsrange().StartSlotEndSlot().StartSlotEndSlot(1, 1).StartSlotEndSlot(1, 1).Build()
	s.ClusterFailover().Force().Build()
	s.ClusterFailover().Takeover().Build()
	s.ClusterFailover().Build()
	s.ClusterFlushslots().Build()
	s.ClusterForget().NodeId("1").Build()
	s.ClusterGetkeysinslot().Slot(1).Count(1).Build()
	s.ClusterInfo().Build()
	s.ClusterKeyslot().Key("1").Build()
	s.ClusterLinks().Build()
	s.ClusterMeet().Ip("1").Port(1).ClusterBusPort(1).Build()
	s.ClusterMeet().Ip("1").Port(1).Build()
	s.ClusterMyid().Build()
	s.ClusterMyshardid().Build()
	s.ClusterNodes().Build()
	s.ClusterReplicas().NodeId("1").Build()
	s.ClusterReplicate().NodeId("1").Build()
	s.ClusterReset().Hard().Build()
	s.ClusterReset().Soft().Build()
	s.ClusterReset().Build()
	s.ClusterSaveconfig().Build()
	s.ClusterSetConfigEpoch().ConfigEpoch(1).Build()
	s.ClusterSetslot().Slot(1).Importing().NodeId("1").Build()
	s.ClusterSetslot().Slot(1).Importing().Build()
	s.ClusterSetslot().Slot(1).Migrating().NodeId("1").Build()
	s.ClusterSetslot().Slot(1).Migrating().Build()
	s.ClusterSetslot().Slot(1).Stable().NodeId("1").Build()
	s.ClusterSetslot().Slot(1).Stable().Build()
	s.ClusterSetslot().Slot(1).Node().NodeId("1").Build()
	s.ClusterSetslot().Slot(1).Node().Build()
	s.ClusterShards().Build()
	s.ClusterSlaves().NodeId("1").Build()
	s.ClusterSlots().Build()
	s.Readonly().Build()
	s.Readwrite().Build()
}

func TestCommand_InitSlot_cluster(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { cluster0(s) })
}

func TestCommand_NoSlot_cluster(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { cluster0(s) })
}
