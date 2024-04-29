import { Menu as AntMenu } from "antd"
import { useMenu } from "@refinedev/core"
import { useThemedLayoutContext } from "@refinedev/antd"

interface MenuProps {
    meta: Record<string, unknown>
}
export default function Menu({ meta }: MenuProps) {
    const { menuItems, selectedKey, defaultOpenKeys } = useMenu({ meta })
    const { setMobileSiderOpen } = useThemedLayoutContext()
    console.log(menuItems)

    const items = menuItems.map(item => ({
        key: item.key,
        icon: meta?.icon ?? <IconAntDesignUnorderedListOutlined />,
        label: meta?.label ?? item.label,
    }))

    return (
        <AntMenu
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
            items={items}
            onClick={() => setMobileSiderOpen(false)}
        />
    )
}
