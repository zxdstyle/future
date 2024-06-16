import { Outlet, useLocation, useNavigate } from "react-router-dom"
import type { MenuProps, ThemeConfig } from "antd"
import { ConfigProvider, Menu } from "antd"
import { useContext, useEffect } from "react"
import { NavbarContext } from "@/components/layout/Navbar/context.tsx"

const theme: ThemeConfig = {
    components: {
        Menu: {
            itemHeight: 32,
            itemSelectedBg: "rgba(171, 172, 186, 0.12)",
            itemSelectedColor: "#fff",
        },
    },
}

export function Component() {
    const location = useLocation()

    const items: MenuProps["items"] = [
        {
            type: "group",
            label: "通用",
            children: [
                { key: "/settings/general", label: "通用", icon: <IconTablerSettings /> },
                { key: "/settings/account", label: "账户", icon: <IconPhUserBold /> },
                { key: "/settings/appearance", label: "外观", icon: <IconPhPaintBrushBold /> },
                { key: "/settings/storage", label: "存储", icon: <IconPhFloppyDisk /> },
            ],
        },
        {
            type: "group",
            label: "系统",
            children: [
                { key: "/settings/about", label: "关于", icon: <IconTablerUfo /> },
                { key: "/settings/release", label: "更新日志", icon: <IconTablerAlignBoxLeftStretch /> },
            ],
        },
    ]

    const navigate = useNavigate()
    const handleClick: MenuProps["onClick"] = ({ key }) => {
        navigate(key)
    }

    const { setLeftTools } = useContext(NavbarContext)
    useEffect(() => {
        setLeftTools(["navigator"])
    }, [])
    return (
        <div className="flex w-full h-full">
            <nav className="w-52 h-full pl-3">
                <ConfigProvider theme={theme}>
                    <Menu items={items} className="h-full pr-3" selectedKeys={[location.pathname]} onClick={handleClick} />
                </ConfigProvider>
            </nav>

            <Outlet />
        </div>
    )
}
