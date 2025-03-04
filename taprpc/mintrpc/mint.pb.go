// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: mintrpc/mint.proto

package mintrpc

import (
	taprpc "github.com/lightninglabs/taproot-assets/taprpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BatchState int32

const (
	BatchState_BATCH_STATE_UNKNOWN            BatchState = 0
	BatchState_BATCH_STATE_PEDNING            BatchState = 1
	BatchState_BATCH_STATE_FROZEN             BatchState = 2
	BatchState_BATCH_STATE_COMMITTED          BatchState = 3
	BatchState_BATCH_STATE_BROADCAST          BatchState = 4
	BatchState_BATCH_STATE_CONFIRMED          BatchState = 5
	BatchState_BATCH_STATE_FINALIZED          BatchState = 6
	BatchState_BATCH_STATE_SEEDLING_CANCELLED BatchState = 7
	BatchState_BATCH_STATE_SPROUT_CANCELLED   BatchState = 8
)

// Enum value maps for BatchState.
var (
	BatchState_name = map[int32]string{
		0: "BATCH_STATE_UNKNOWN",
		1: "BATCH_STATE_PEDNING",
		2: "BATCH_STATE_FROZEN",
		3: "BATCH_STATE_COMMITTED",
		4: "BATCH_STATE_BROADCAST",
		5: "BATCH_STATE_CONFIRMED",
		6: "BATCH_STATE_FINALIZED",
		7: "BATCH_STATE_SEEDLING_CANCELLED",
		8: "BATCH_STATE_SPROUT_CANCELLED",
	}
	BatchState_value = map[string]int32{
		"BATCH_STATE_UNKNOWN":            0,
		"BATCH_STATE_PEDNING":            1,
		"BATCH_STATE_FROZEN":             2,
		"BATCH_STATE_COMMITTED":          3,
		"BATCH_STATE_BROADCAST":          4,
		"BATCH_STATE_CONFIRMED":          5,
		"BATCH_STATE_FINALIZED":          6,
		"BATCH_STATE_SEEDLING_CANCELLED": 7,
		"BATCH_STATE_SPROUT_CANCELLED":   8,
	}
)

func (x BatchState) Enum() *BatchState {
	p := new(BatchState)
	*p = x
	return p
}

func (x BatchState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BatchState) Descriptor() protoreflect.EnumDescriptor {
	return file_mintrpc_mint_proto_enumTypes[0].Descriptor()
}

func (BatchState) Type() protoreflect.EnumType {
	return &file_mintrpc_mint_proto_enumTypes[0]
}

func (x BatchState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BatchState.Descriptor instead.
func (BatchState) EnumDescriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{0}
}

type MintAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the asset to be created.
	AssetType taprpc.AssetType `protobuf:"varint,1,opt,name=asset_type,json=assetType,proto3,enum=taprpc.AssetType" json:"asset_type,omitempty"`
	// The name, or "tag" of the asset. This will affect the final asset ID.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// A blob that resents metadata related to the asset. This will affect the
	// final asset ID.
	AssetMeta *taprpc.AssetMeta `protobuf:"bytes,3,opt,name=asset_meta,json=assetMeta,proto3" json:"asset_meta,omitempty"`
	// The total amount of units of the new asset that should be created. If the
	// AssetType is Collectible, then this field cannot be set.
	Amount uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	// The specific group key this asset should be minted with.
	GroupKey []byte `protobuf:"bytes,5,opt,name=group_key,json=groupKey,proto3" json:"group_key,omitempty"`
	// The name of the asset in the batch that will anchor a new asset group.
	// This asset will be minted with the same group key as the anchor asset.
	GroupAnchor string `protobuf:"bytes,6,opt,name=group_anchor,json=groupAnchor,proto3" json:"group_anchor,omitempty"`
}

func (x *MintAsset) Reset() {
	*x = MintAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintrpc_mint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintAsset) ProtoMessage() {}

func (x *MintAsset) ProtoReflect() protoreflect.Message {
	mi := &file_mintrpc_mint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintAsset.ProtoReflect.Descriptor instead.
func (*MintAsset) Descriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{0}
}

func (x *MintAsset) GetAssetType() taprpc.AssetType {
	if x != nil {
		return x.AssetType
	}
	return taprpc.AssetType(0)
}

func (x *MintAsset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MintAsset) GetAssetMeta() *taprpc.AssetMeta {
	if x != nil {
		return x.AssetMeta
	}
	return nil
}

func (x *MintAsset) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *MintAsset) GetGroupKey() []byte {
	if x != nil {
		return x.GroupKey
	}
	return nil
}

func (x *MintAsset) GetGroupAnchor() string {
	if x != nil {
		return x.GroupAnchor
	}
	return ""
}

type MintAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The asset to be minted.
	Asset *MintAsset `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	// If true, then the asset will be created with a group key, which allows for
	// future asset issuance.
	EnableEmission bool `protobuf:"varint,2,opt,name=enable_emission,json=enableEmission,proto3" json:"enable_emission,omitempty"`
}

func (x *MintAssetRequest) Reset() {
	*x = MintAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintrpc_mint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintAssetRequest) ProtoMessage() {}

func (x *MintAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mintrpc_mint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintAssetRequest.ProtoReflect.Descriptor instead.
func (*MintAssetRequest) Descriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{1}
}

func (x *MintAssetRequest) GetAsset() *MintAsset {
	if x != nil {
		return x.Asset
	}
	return nil
}

func (x *MintAssetRequest) GetEnableEmission() bool {
	if x != nil {
		return x.EnableEmission
	}
	return false
}

type MintAssetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A public key serialized in compressed format that can be used to uniquely
	// identify a pending minting batch. Responses that share the same key will be
	// batched into the same minting transaction.
	BatchKey []byte `protobuf:"bytes,1,opt,name=batch_key,json=batchKey,proto3" json:"batch_key,omitempty"`
}

func (x *MintAssetResponse) Reset() {
	*x = MintAssetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintrpc_mint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintAssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintAssetResponse) ProtoMessage() {}

func (x *MintAssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mintrpc_mint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintAssetResponse.ProtoReflect.Descriptor instead.
func (*MintAssetResponse) Descriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{2}
}

func (x *MintAssetResponse) GetBatchKey() []byte {
	if x != nil {
		return x.BatchKey
	}
	return nil
}

type MintingBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The internal public key of the batch.
	BatchKey []byte `protobuf:"bytes,1,opt,name=batch_key,json=batchKey,proto3" json:"batch_key,omitempty"`
	// The assets that are part of the batch.
	Assets []*MintAsset `protobuf:"bytes,2,rep,name=assets,proto3" json:"assets,omitempty"`
	// The state of the batch.
	State BatchState `protobuf:"varint,3,opt,name=state,proto3,enum=mintrpc.BatchState" json:"state,omitempty"`
}

func (x *MintingBatch) Reset() {
	*x = MintingBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintrpc_mint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintingBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintingBatch) ProtoMessage() {}

func (x *MintingBatch) ProtoReflect() protoreflect.Message {
	mi := &file_mintrpc_mint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintingBatch.ProtoReflect.Descriptor instead.
func (*MintingBatch) Descriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{3}
}

func (x *MintingBatch) GetBatchKey() []byte {
	if x != nil {
		return x.BatchKey
	}
	return nil
}

func (x *MintingBatch) GetAssets() []*MintAsset {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *MintingBatch) GetState() BatchState {
	if x != nil {
		return x.State
	}
	return BatchState_BATCH_STATE_UNKNOWN
}

type FinalizeBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinalizeBatchRequest) Reset() {
	*x = FinalizeBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintrpc_mint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeBatchRequest) ProtoMessage() {}

func (x *FinalizeBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mintrpc_mint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeBatchRequest.ProtoReflect.Descriptor instead.
func (*FinalizeBatchRequest) Descriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{4}
}

type FinalizeBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The internal public key of the batch.
	BatchKey []byte `protobuf:"bytes,1,opt,name=batch_key,json=batchKey,proto3" json:"batch_key,omitempty"`
}

func (x *FinalizeBatchResponse) Reset() {
	*x = FinalizeBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintrpc_mint_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeBatchResponse) ProtoMessage() {}

func (x *FinalizeBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mintrpc_mint_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeBatchResponse.ProtoReflect.Descriptor instead.
func (*FinalizeBatchResponse) Descriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{5}
}

func (x *FinalizeBatchResponse) GetBatchKey() []byte {
	if x != nil {
		return x.BatchKey
	}
	return nil
}

type CancelBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelBatchRequest) Reset() {
	*x = CancelBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintrpc_mint_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBatchRequest) ProtoMessage() {}

func (x *CancelBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mintrpc_mint_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBatchRequest.ProtoReflect.Descriptor instead.
func (*CancelBatchRequest) Descriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{6}
}

type CancelBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The internal public key of the batch.
	BatchKey []byte `protobuf:"bytes,1,opt,name=batch_key,json=batchKey,proto3" json:"batch_key,omitempty"`
}

func (x *CancelBatchResponse) Reset() {
	*x = CancelBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintrpc_mint_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBatchResponse) ProtoMessage() {}

func (x *CancelBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mintrpc_mint_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBatchResponse.ProtoReflect.Descriptor instead.
func (*CancelBatchResponse) Descriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{7}
}

func (x *CancelBatchResponse) GetBatchKey() []byte {
	if x != nil {
		return x.BatchKey
	}
	return nil
}

type ListBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The optional batch key of the batch to list. When using REST this field
	// must be encoded as base64url.
	BatchKey []byte `protobuf:"bytes,1,opt,name=batch_key,json=batchKey,proto3" json:"batch_key,omitempty"`
}

func (x *ListBatchRequest) Reset() {
	*x = ListBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintrpc_mint_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBatchRequest) ProtoMessage() {}

func (x *ListBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mintrpc_mint_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBatchRequest.ProtoReflect.Descriptor instead.
func (*ListBatchRequest) Descriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{8}
}

func (x *ListBatchRequest) GetBatchKey() []byte {
	if x != nil {
		return x.BatchKey
	}
	return nil
}

type ListBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batches []*MintingBatch `protobuf:"bytes,1,rep,name=batches,proto3" json:"batches,omitempty"`
}

func (x *ListBatchResponse) Reset() {
	*x = ListBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintrpc_mint_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBatchResponse) ProtoMessage() {}

func (x *ListBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mintrpc_mint_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBatchResponse.ProtoReflect.Descriptor instead.
func (*ListBatchResponse) Descriptor() ([]byte, []int) {
	return file_mintrpc_mint_proto_rawDescGZIP(), []int{9}
}

func (x *ListBatchResponse) GetBatches() []*MintingBatch {
	if x != nil {
		return x.Batches
	}
	return nil
}

var File_mintrpc_mint_proto protoreflect.FileDescriptor

var file_mintrpc_mint_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x1a, 0x13, 0x74,
	0x61, 0x70, 0x72, 0x6f, 0x6f, 0x74, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x09, 0x4d, 0x69, 0x6e, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x12, 0x30, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x70, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x70,
	0x72, 0x70, 0x63, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x09, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72,
	0x22, 0x65, 0x0a, 0x10, 0x4d, 0x69, 0x6e, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x69,
	0x6e, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x45,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x11, 0x4d, 0x69, 0x6e, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x4d, 0x69,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x06, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x16,
	0x0a, 0x14, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x14, 0x0a, 0x12,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x32, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x2f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x44, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2a, 0x88, 0x02,
	0x0a, 0x0a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x13,
	0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x45, 0x44, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x16,
	0x0a, 0x12, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x52,
	0x4f, 0x5a, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15,
	0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44,
	0x10, 0x06, 0x12, 0x22, 0x0a, 0x1e, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x53, 0x45, 0x45, 0x44, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x12, 0x20, 0x0a, 0x1c, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x50, 0x52, 0x4f, 0x55, 0x54, 0x5f, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x08, 0x32, 0xaa, 0x02, 0x0a, 0x04, 0x4d, 0x69, 0x6e,
	0x74, 0x12, 0x42, 0x0a, 0x09, 0x4d, 0x69, 0x6e, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x19,
	0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x69, 0x6e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x19,
	0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x69, 0x6e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x6f, 0x74, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x74, 0x61, 0x70, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mintrpc_mint_proto_rawDescOnce sync.Once
	file_mintrpc_mint_proto_rawDescData = file_mintrpc_mint_proto_rawDesc
)

func file_mintrpc_mint_proto_rawDescGZIP() []byte {
	file_mintrpc_mint_proto_rawDescOnce.Do(func() {
		file_mintrpc_mint_proto_rawDescData = protoimpl.X.CompressGZIP(file_mintrpc_mint_proto_rawDescData)
	})
	return file_mintrpc_mint_proto_rawDescData
}

var file_mintrpc_mint_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mintrpc_mint_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_mintrpc_mint_proto_goTypes = []interface{}{
	(BatchState)(0),               // 0: mintrpc.BatchState
	(*MintAsset)(nil),             // 1: mintrpc.MintAsset
	(*MintAssetRequest)(nil),      // 2: mintrpc.MintAssetRequest
	(*MintAssetResponse)(nil),     // 3: mintrpc.MintAssetResponse
	(*MintingBatch)(nil),          // 4: mintrpc.MintingBatch
	(*FinalizeBatchRequest)(nil),  // 5: mintrpc.FinalizeBatchRequest
	(*FinalizeBatchResponse)(nil), // 6: mintrpc.FinalizeBatchResponse
	(*CancelBatchRequest)(nil),    // 7: mintrpc.CancelBatchRequest
	(*CancelBatchResponse)(nil),   // 8: mintrpc.CancelBatchResponse
	(*ListBatchRequest)(nil),      // 9: mintrpc.ListBatchRequest
	(*ListBatchResponse)(nil),     // 10: mintrpc.ListBatchResponse
	(taprpc.AssetType)(0),         // 11: taprpc.AssetType
	(*taprpc.AssetMeta)(nil),      // 12: taprpc.AssetMeta
}
var file_mintrpc_mint_proto_depIdxs = []int32{
	11, // 0: mintrpc.MintAsset.asset_type:type_name -> taprpc.AssetType
	12, // 1: mintrpc.MintAsset.asset_meta:type_name -> taprpc.AssetMeta
	1,  // 2: mintrpc.MintAssetRequest.asset:type_name -> mintrpc.MintAsset
	1,  // 3: mintrpc.MintingBatch.assets:type_name -> mintrpc.MintAsset
	0,  // 4: mintrpc.MintingBatch.state:type_name -> mintrpc.BatchState
	4,  // 5: mintrpc.ListBatchResponse.batches:type_name -> mintrpc.MintingBatch
	2,  // 6: mintrpc.Mint.MintAsset:input_type -> mintrpc.MintAssetRequest
	5,  // 7: mintrpc.Mint.FinalizeBatch:input_type -> mintrpc.FinalizeBatchRequest
	7,  // 8: mintrpc.Mint.CancelBatch:input_type -> mintrpc.CancelBatchRequest
	9,  // 9: mintrpc.Mint.ListBatches:input_type -> mintrpc.ListBatchRequest
	3,  // 10: mintrpc.Mint.MintAsset:output_type -> mintrpc.MintAssetResponse
	6,  // 11: mintrpc.Mint.FinalizeBatch:output_type -> mintrpc.FinalizeBatchResponse
	8,  // 12: mintrpc.Mint.CancelBatch:output_type -> mintrpc.CancelBatchResponse
	10, // 13: mintrpc.Mint.ListBatches:output_type -> mintrpc.ListBatchResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_mintrpc_mint_proto_init() }
func file_mintrpc_mint_proto_init() {
	if File_mintrpc_mint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mintrpc_mint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MintAsset); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mintrpc_mint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MintAssetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mintrpc_mint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MintAssetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mintrpc_mint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MintingBatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mintrpc_mint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeBatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mintrpc_mint_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeBatchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mintrpc_mint_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelBatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mintrpc_mint_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelBatchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mintrpc_mint_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mintrpc_mint_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBatchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mintrpc_mint_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mintrpc_mint_proto_goTypes,
		DependencyIndexes: file_mintrpc_mint_proto_depIdxs,
		EnumInfos:         file_mintrpc_mint_proto_enumTypes,
		MessageInfos:      file_mintrpc_mint_proto_msgTypes,
	}.Build()
	File_mintrpc_mint_proto = out.File
	file_mintrpc_mint_proto_rawDesc = nil
	file_mintrpc_mint_proto_goTypes = nil
	file_mintrpc_mint_proto_depIdxs = nil
}
