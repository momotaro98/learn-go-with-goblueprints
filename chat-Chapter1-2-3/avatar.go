package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarURL is an error which is thrown
// when Avatar instance can't return the URL.
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatar is a type which represents image of user profile.
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返す。
	// 問題が発生した場合にはエラーを返す。特に、URLを取得できなかった
	// 場合にはErrNoAvatarURLを返す。
	GetAvatarURL(u ChatUser) (string, error)
}

type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

var UseGravatarAvatar GravatarAvatar

func (_ GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (a FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	filepathStrSlice := []string{"avatars", "build/avatars"}
	for _, filepathStr := range filepathStrSlice {
		if url, err := a.getAvatarURL(u, filepathStr); err == nil {
			return url, err
		}
	}
	return "", ErrNoAvatarURL
}

func (_ FileSystemAvatar) getAvatarURL(u ChatUser, filepathStr string) (string, error) {
	if files, err := ioutil.ReadDir(filepathStr); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
