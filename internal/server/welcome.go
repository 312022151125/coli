// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package server

import (
	"fmt"
	"os"
	"strings"

	versionPkg "github.com/312022151125/coli/internal/version"
	"github.com/charmbracelet/lipgloss"
)

const (
	// GreetingBanner 是控制台横幅
	GreetingBanner = `
███████╗     ██████╗    ██╗  ██╗     ██████╗ 
██╔════╝    ██╔════╝    ██║  ██║    ██╔═████╗
█████╗      ██║         ███████║    ██║██╔██║
██╔══╝      ██║         ██╔══██║    ████╔╝██║
███████╗    ╚██████╗    ██║  ██║    ╚██████╔╝
╚══════╝     ╚═════╝    ╚═╝  ╚═╝     ╚═════╝ 
                                             
`
)

var (
	// 信息样式（每行）
	infoStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(lipgloss.AdaptiveColor{
			Light: "236", Dark: "252",
		})

	// 标题样式
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{
			Light: "#4338ca", Dark: "#f7b457ff",
		})

	// 高亮样式
	highlight = lipgloss.NewStyle().
			Bold(false).
			Italic(true).
			Foreground(lipgloss.AdaptiveColor{
			Light: "#7c3aed", Dark: "#53b7f5ff",
		})

	// 外框
	boxStyle = lipgloss.NewStyle().
			Bold(true).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#fb5151ff")).
			Padding(1, 1).
			Margin(1, 1)
)

// PrintGreetings 使用 lipgloss 输出欢迎信息
func PrintGreetings(port string) {
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		infoStyle.Render(
			"📦 "+titleStyle.Render("Version")+": "+highlight.Render(versionPkg.Version),
		),
		infoStyle.Render("🎈 "+titleStyle.Render("Port")+": "+highlight.Render(port)),
		infoStyle.Render(
			"🌐 "+titleStyle.Render("Website")+": "+highlight.Render("https://coli.dev/"),
		),
		infoStyle.Render(
			"📦 "+titleStyle.Render("Source")+": "+highlight.Render(versionPkg.SourceURL),
		),
	)

	full := boxStyle.Render(content)

	if _, err := fmt.Fprintln(os.Stdout, full); err != nil {
		fmt.Fprintf(os.Stderr, "failed to print greetings: %v\n", err)
	}
}

func gradientBanner(banner string) string {
	lines := strings.Split(banner, "\n")
	var rendered []string

	colors := []string{
		"#FF7F7F", // 珊瑚红
		"#FFB347", // 桃橙色
		"#FFEB9C", // 金黄色
		"#B8E6B8", // 薄荷绿
		"#87CEEB", // 天空蓝
		"#DDA0DD", // 梅花紫
		"#F0E68C", // 卡其色
	}

	for i, line := range lines {
		color := lipgloss.Color(colors[i%len(colors)])
		style := lipgloss.NewStyle().Foreground(color)
		rendered = append(rendered, style.Render(line))
	}
	return lipgloss.JoinVertical(lipgloss.Left, rendered...)
}
