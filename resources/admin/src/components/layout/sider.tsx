import type { CSSProperties } from "react"
import React, { useContext } from "react"
import {
    useActiveAuthProvider,
    useIsExistAuthentication,
    useLink,
    useLogout,
    useMenu,
    useRefineContext,
    useRouterContext,
    useRouterType,
    useTitle,
    useTranslate,
    useWarnAboutChange,
} from "@refinedev/core"
import { useThemedLayoutContext } from "@refinedev/antd"

import { Button, Drawer, Grid, Layout, Menu, theme } from "antd"
import type { RefineThemedLayoutV2SiderProps } from "@refinedev/antd"
// @ts-expect-error
import { ThemedTitleV2 } from "./title"
import { ColorModeContext } from "@/contexts/color-mode"

const drawerButtonStyles: CSSProperties = {
    borderTopLeftRadius: 0,
    borderBottomLeftRadius: 0,
    position: "fixed",
    top: 64,
    zIndex: 999,
}

export const ThemedSiderV2: React.FC<RefineThemedLayoutV2SiderProps> = ({
    Title: TitleFromProps,
    render,
    meta,
    fixed,
    activeItemDisabled = false,
}) => {
    const { token } = theme.useToken()
    const { siderCollapsed, setSiderCollapsed, mobileSiderOpen, setMobileSiderOpen } = useThemedLayoutContext()

    const isExistAuthentication = useIsExistAuthentication()
    const routerType = useRouterType()
    const NewLink = useLink()
    const { warnWhen, setWarnWhen } = useWarnAboutChange()
    const { Link: LegacyLink } = useRouterContext()
    const Link = routerType === "legacy" ? LegacyLink : NewLink
    const TitleFromContext = useTitle()
    const translate = useTranslate()
    const { menuItems, selectedKey, defaultOpenKeys } = useMenu({ meta })
    const breakpoint = Grid.useBreakpoint()
    const { hasDashboard } = useRefineContext()
    const authProvider = useActiveAuthProvider()
    const { mutate: mutateLogout } = useLogout({
        v3LegacyAuthProviderCompatible: Boolean(authProvider?.isLegacy),
    })
    console.log(menuItems)
    const { mode } = useContext(ColorModeContext)

    const isMobile = typeof breakpoint.lg === "undefined" ? false : !breakpoint.lg

    const RenderToTitle = TitleFromProps ?? TitleFromContext ?? ThemedTitleV2

    const renderTreeView = (tree: TreeMenuItem[], selectedKey?: string) => {
        return tree.map((item: TreeMenuItem) => {
            const { key, name, children, meta, route } = item

            if (children.length > 0) {
                return (
                    <CanAccess
                        key={item.key}
                        resource={name.toLowerCase()}
                        action="list"
                        params={{
                            resource: item,
                        }}
                    >
                        <Menu.SubMenu
                            key={item.key}
                            icon={meta?.icon ?? <IconAntDesignUnorderedListOutlined />}
                            title={meta?.label ?? name}
                        >
                            {renderTreeView(children, selectedKey)}
                        </Menu.SubMenu>
                    </CanAccess>
                )
            }
            const isSelected = key === selectedKey
            const isRoute = !(meta?.parent !== undefined && children.length === 0)

            const linkStyle: React.CSSProperties = activeItemDisabled && isSelected ? { pointerEvents: "none" } : {}

            return (
                <CanAccess
                    key={item.key}
                    resource={name.toLowerCase()}
                    action="list"
                    params={{
                        resource: item,
                    }}
                >
                    <Menu.Item
                        key={item.key}
                        icon={meta?.icon ?? (isRoute && <IconAntDesignUnorderedListOutlined />)}
                        style={linkStyle}
                    >
                        <Link to={route ?? ""} style={linkStyle}>
                            {meta?.label ?? name}
                        </Link>
                        {!siderCollapsed && isSelected && <div className="ant-menu-tree-arrow" />}
                    </Menu.Item>
                </CanAccess>
            )
        })
    }

    const handleLogout = () => {
        if (warnWhen) {
            const confirm = window.confirm(
                translate("warnWhenUnsavedChanges", "Are you sure you want to leave? You have unsaved changes."),
            )

            if (confirm) {
                setWarnWhen(false)
                mutateLogout()
            }
        }
        else {
            mutateLogout()
        }
    }

    const logout = isExistAuthentication && (
        <Menu.Item key="logout" onClick={() => handleLogout()} icon={<IconAntDesignLogoutOutlined />}>
            {translate("buttons.logout", "Logout")}
        </Menu.Item>
    )

    const dashboard = hasDashboard
        ? (
            <Menu.Item key="dashboard" icon={<IconAntDesignDashboardOutlined />}>
                <Link to="/">{translate("dashboard.title", "Dashboard")}</Link>
                {!siderCollapsed && selectedKey === "/" && <div className="ant-menu-tree-arrow" />}
            </Menu.Item>
        )
        : null

    const items = renderTreeView(menuItems, selectedKey)

    const renderSider = () => {
        if (render) {
            return render({
                dashboard,
                items,
                logout,
                collapsed: siderCollapsed,
            })
        }
        return (
            <>
                {dashboard}
                {items}
                {logout}
            </>
        )
    }

    const renderMenu = () => {
        return (
            <Menu
                selectedKeys={selectedKey ? [selectedKey] : []}
                defaultOpenKeys={defaultOpenKeys}
                mode="inline"
                style={{
                    paddingTop: "8px",
                    padding: "0 8px",
                    border: "none",
                    overflow: "auto",
                    height: "calc(100% - 72px)",
                    transition: "all 0.2s",
                }}
                onClick={() => {
                    setMobileSiderOpen(false)
                }}
            >
                {renderSider()}
            </Menu>
        )
    }

    const renderDrawerSider = () => {
        return (
            <>
                <Drawer
                    open={mobileSiderOpen}
                    onClose={() => setMobileSiderOpen(false)}
                    placement="left"
                    closable={false}
                    width={200}
                    styles={{ body: { padding: 0 } }}
                    maskClosable={true}
                >
                    <Layout>
                        <Layout.Sider
                            style={{
                                height: "100vh",
                                backgroundColor: token.colorBgContainer,
                                borderRight: `1px solid ${token.colorBgElevated}`,
                            }}
                        >
                            <div
                                style={{
                                    width: "200px",
                                    padding: "0 16px",
                                    display: "flex",
                                    justifyContent: "flex-start",
                                    alignItems: "center",
                                    height: "64px",
                                    backgroundColor: token.colorBgElevated,
                                }}
                            >
                                <RenderToTitle collapsed={false} />
                            </div>
                            {renderMenu()}
                        </Layout.Sider>
                    </Layout>
                </Drawer>
                <Button
                    style={drawerButtonStyles}
                    size="large"
                    onClick={() => setMobileSiderOpen(true)}
                    icon={<IconAntDesignBarsOutlined />}
                >
                </Button>
            </>
        )
    }

    if (isMobile)
        return renderDrawerSider()

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
                {renderMenu()}
            </Layout.Sider>
        </>
    )
}
