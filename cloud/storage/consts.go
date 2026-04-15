package storage

import (
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos/enum"
)

// 访问权限类型
const (
	ACLPrivate         = "private"
	ACLPublicRead      = "public-read"
	ACLPublicReadWrite = "public-read-write"
	ACLDefault         = "default"
)

// 存储类型
const (
	StorageClassStandard        = "STANDARD"
	StorageClassIa              = "IA"
	StorageClassArchive         = "ARCHIVE"
	StorageClassColdArchive     = "COLD_ARCHIVE"
	StorageClassDeepColdArchive = "DEEP_COLD_ARCHIVE" // StorageClassDeepColdArchive
)

func ToAliyunACL(acl string) oss.ObjectACLType {
	switch acl {
	case ACLPrivate:
		return oss.ObjectACLPrivate
	case ACLPublicRead:
		return oss.ObjectACLPublicRead
	case ACLPublicReadWrite:
		return oss.ObjectACLPublicReadWrite
	case ACLDefault:
		return oss.ObjectACLDefault
	}
	return oss.ObjectACLDefault
}

func ToVolcACL(acl string) enum.ACLType {
	switch acl {
	case ACLPrivate:
		return enum.ACLPrivate
	case ACLPublicRead:
		return enum.ACLPublicRead
	case ACLPublicReadWrite:
		return enum.ACLPublicReadWrite
	case ACLDefault:
		return enum.ACLDefault
	}
	return enum.ACLDefault
}

func ToAliyunStorageClass(storageClass string) oss.StorageClassType {
	switch storageClass {
	case StorageClassStandard:
		return oss.StorageClassStandard
	case StorageClassIa:
		return oss.StorageClassIA
	case StorageClassArchive:
		return oss.StorageClassArchive
	case StorageClassColdArchive:
		return oss.StorageClassColdArchive
	case StorageClassDeepColdArchive:
		return oss.StorageClassDeepColdArchive
	}
	return ""
}

func ToVolcStorageClass(storageClass string) enum.StorageClassType {
	switch storageClass {
	case StorageClassStandard:
		return enum.StorageClassStandard
	case StorageClassIa:
		return enum.StorageClassIa
	case StorageClassArchive:
		return enum.StorageClassArchive
	case StorageClassColdArchive:
		return enum.StorageClassColdArchive
	case StorageClassDeepColdArchive:
		return enum.StorageClassDeepColdArchive
	}
	return ""
}
