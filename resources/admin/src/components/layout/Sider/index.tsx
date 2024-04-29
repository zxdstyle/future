import type { RefineThemedLayoutV2SiderProps } from "@refinedev/antd"
import { useThemedLayoutContext } from "@refinedev/antd"
import { Button, Layout, theme } from "antd"
import React, { useContext } from "react"
import { useTitle } from "@refinedev/core"
import Menu from "./Menu"
import { ColorModeContext } from "@/contexts/color-mode"
import { ThemedTitleV2 } from "@/components/layout"

export default function Sider({ fixed, meta, Title: TitleFromProps }: RefineThemedLayoutV2SiderProps) {
    const { siderCollapsed, setSiderCollapsed } = useThemedLayoutContext()
    const { token } = theme.useToken()
    const { mode } = useContext(ColorModeContext)

    const siderStyles: React.CSSProperties = {
        // backgroundColor: token.colorBgContainer,
        // borderRight: `1px solid ${token.colorBgElevated}`,
    }

    if (fixed) {
        siderStyles.position = "fixed"
        siderStyles.top = 0
        siderStyles.height = "100vh"
        siderStyles.zIndex = 999
    }

    const TitleFromContext = useTitle()
    const RenderToTitle = TitleFromProps ?? TitleFromContext ?? ThemedTitleV2

    return (
        <>
            {fixed && (
                <div
                    style={{
                        width: siderCollapsed ? "80px" : "200px",
                        transition: "all 0.2s",
                    }}
                />
            )}
            <Layout.Sider
                style={siderStyles}
                collapsible
                collapsed={siderCollapsed}
                onCollapse={(collapsed, type) => {
                    if (type === "clickTrigger")
                        setSiderCollapsed(collapsed)
                }}
                collapsedWidth={80}
                breakpoint="lg"
                theme={mode}
                trigger={(
                    <Button
                        shape="circle"
                        className="absolute top-6 -right-4 shadow border-none"
                        icon={siderCollapsed ? <IconAntDesignRightOutlined /> : <IconAntDesignLeftOutlined />}
                        onClick={() => {}}
                    >
                    </Button>
                )}
            >
                <div
                    style={{
                        width: siderCollapsed ? "80px" : "200px",
                        padding: siderCollapsed ? "0" : "0 16px",
                        display: "flex",
                        justifyContent: siderCollapsed ? "center" : "flex-start",
                        alignItems: "center",
                        height: "64px",
                        backgroundColor: token.colorBgElevated,
                        fontSize: "14px",
                        transition: "all 0.2s",
                    }}
                >
                    <RenderToTitle collapsed={siderCollapsed} />
                </div>

                <Menu meta={meta} />
            </Layout.Sider>
        </>

    )
}
