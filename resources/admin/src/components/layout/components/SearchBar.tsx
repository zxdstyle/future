import type { ThemeConfig } from "antd"
import { ConfigProvider, Input } from "antd"

export default function SearchBar() {
    const theme: ThemeConfig = {
        components: {
            Input: {
                colorBorder: "transparent",
                colorBgContainer: "transparent",
                controlOutlineWidth: 0,
                controlOutline: "transparent",
            },
        },
    }

    return (
        <ConfigProvider theme={theme}>
            <Input
                className="w-full md:w-52 bg-slate-100 rounded-3xl text-sm font-bold !border-none"
                size="large"
                placeholder="Search..."
                prefix={<IconBiSearch className="text-gray-700" />}
            >
            </Input>
        </ConfigProvider>
    )
}
