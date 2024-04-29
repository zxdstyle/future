import React from "react"
import { useLink, useRouterContext, useRouterType } from "@refinedev/core"
import { Space, Typography, theme } from "antd"
import type { RefineLayoutThemedTitleProps } from "@refinedev/antd"

const defaultText = "Future Project"

export const ThemedTitleV2: React.FC<RefineLayoutThemedTitleProps> = ({
    collapsed,
    icon,
    text = defaultText,
    wrapperStyles,
}) => {
    const { token } = theme.useToken()
    const routerType = useRouterType()
    const Link = useLink()
    const { Link: LegacyLink } = useRouterContext()

    const ActiveLink = routerType === "legacy" ? LegacyLink : Link

    return (
        <ActiveLink
            to="/"
            style={{
                display: "inline-block",
                textDecoration: "none",
            }}
        >
            <Space
                style={{
                    display: "flex",
                    alignItems: "center",
                    fontSize: "inherit",
                    ...wrapperStyles,
                }}
            >
                <div
                    style={{
                        height: "24px",
                        width: "24px",
                        color: token.colorPrimary,
                    }}
                >
                    {icon}
                </div>

                {!collapsed && (
                    <Typography.Title
                        style={{
                            fontSize: "inherit",
                            marginBottom: 0,
                            fontWeight: 700,
                        }}
                    >
                        {text}
                    </Typography.Title>
                )}
            </Space>
        </ActiveLink>
    )
}
