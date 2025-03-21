package util

import (
	ethpb "github.com/prysmaticlabs/prysm/v5/proto/prysm/v1alpha1"
)

// ----------------------------------------------------------------------------
// Bellatrix
// ----------------------------------------------------------------------------

// NewBeaconBlockBellatrix creates a beacon block with minimum marshalable fields.
func NewBeaconBlockBellatrix() *ethpb.SignedBeaconBlockBellatrix {
	return HydrateSignedBeaconBlockBellatrix(&ethpb.SignedBeaconBlockBellatrix{})
}

// NewBlindedBeaconBlockBellatrix creates a blinded beacon block with minimum marshalable fields.
func NewBlindedBeaconBlockBellatrix() *ethpb.SignedBlindedBeaconBlockBellatrix {
	return HydrateSignedBlindedBeaconBlockBellatrix(&ethpb.SignedBlindedBeaconBlockBellatrix{})
}

// ----------------------------------------------------------------------------
// Capella
// ----------------------------------------------------------------------------

// NewBeaconBlockCapella creates a beacon block with minimum marshalable fields.
func NewBeaconBlockCapella() *ethpb.SignedBeaconBlockCapella {
	return HydrateSignedBeaconBlockCapella(&ethpb.SignedBeaconBlockCapella{})
}

// NewBlindedBeaconBlockCapella creates a blinded beacon block with minimum marshalable fields.
func NewBlindedBeaconBlockCapella() *ethpb.SignedBlindedBeaconBlockCapella {
	return HydrateSignedBlindedBeaconBlockCapella(&ethpb.SignedBlindedBeaconBlockCapella{})
}

// ----------------------------------------------------------------------------
// Deneb
// ----------------------------------------------------------------------------

// NewBeaconBlockDeneb creates a beacon block with minimum marshalable fields.
func NewBeaconBlockDeneb() *ethpb.SignedBeaconBlockDeneb {
	return HydrateSignedBeaconBlockDeneb(&ethpb.SignedBeaconBlockDeneb{})
}

// NewBeaconBlockContentsDeneb creates a beacon block with minimum marshalable fields.
func NewBeaconBlockContentsDeneb() *ethpb.SignedBeaconBlockContentsDeneb {
	return HydrateSignedBeaconBlockContentsDeneb(&ethpb.SignedBeaconBlockContentsDeneb{})
}

// NewBlindedBeaconBlockDeneb creates a blinded beacon block with minimum marshalable fields.
func NewBlindedBeaconBlockDeneb() *ethpb.SignedBlindedBeaconBlockDeneb {
	return HydrateSignedBlindedBeaconBlockDeneb(&ethpb.SignedBlindedBeaconBlockDeneb{})
}

// ----------------------------------------------------------------------------
// Electra
// ----------------------------------------------------------------------------

// NewBeaconBlockElectra creates a beacon block with minimum marshalable fields.
func NewBeaconBlockElectra() *ethpb.SignedBeaconBlockElectra {
	return HydrateSignedBeaconBlockElectra(&ethpb.SignedBeaconBlockElectra{})
}

// NewBeaconBlockContentsElectra creates a beacon block with minimum marshalable fields.
func NewBeaconBlockContentsElectra() *ethpb.SignedBeaconBlockContentsElectra {
	return HydrateSignedBeaconBlockContentsElectra(&ethpb.SignedBeaconBlockContentsElectra{})
}

// NewBlindedBeaconBlockElectra creates a blinded beacon block with minimum marshalable fields.
func NewBlindedBeaconBlockElectra() *ethpb.SignedBlindedBeaconBlockElectra {
	return HydrateSignedBlindedBeaconBlockElectra(&ethpb.SignedBlindedBeaconBlockElectra{})
}

// ----------------------------------------------------------------------------
// Fulu
// ----------------------------------------------------------------------------

// NewBeaconBlockFulu creates a beacon block with minimum marshalable fields.
func NewBeaconBlockFulu() *ethpb.SignedBeaconBlockFulu {
	return HydrateSignedBeaconBlockFulu(&ethpb.SignedBeaconBlockFulu{})
}

// NewBeaconBlockContentsFulu creates a beacon block with minimum marshalable fields.
func NewBeaconBlockContentsFulu() *ethpb.SignedBeaconBlockContentsFulu {
	return HydrateSignedBeaconBlockContentsFulu(&ethpb.SignedBeaconBlockContentsFulu{})
}

// NewBlindedBeaconBlockFulu creates a blinded beacon block with minimum marshalable fields.
func NewBlindedBeaconBlockFulu() *ethpb.SignedBlindedBeaconBlockFulu {
	return HydrateSignedBlindedBeaconBlockFulu(&ethpb.SignedBlindedBeaconBlockFulu{})
}

// ----------------------------------------------------------------------------
// Eip7805
// ----------------------------------------------------------------------------

// NewBeaconBlockEip7805 creates a beacon block with minimum marshalable fields.
func NewBeaconBlockEip7805() *ethpb.SignedBeaconBlockEip7805 {
	return HydrateSignedBeaconBlockEip7805(&ethpb.SignedBeaconBlockEip7805{})
}

// NewBeaconBlockContentsEip7805 creates a beacon block with minimum marshalable fields.
func NewBeaconBlockContentsEip7805() *ethpb.SignedBeaconBlockContentsEip7805 {
	return HydrateSignedBeaconBlockContentsEip7805(&ethpb.SignedBeaconBlockContentsEip7805{})
}

// NewBlindedBeaconBlockEip7805 creates a blinded beacon block with minimum marshalable fields.
func NewBlindedBeaconBlockEip7805() *ethpb.SignedBlindedBeaconBlockEip7805 {
	return HydrateSignedBlindedBeaconBlockEip7805(&ethpb.SignedBlindedBeaconBlockEip7805{})
}
