import type { SelectProps, ThemeConfig } from "antd"
import { Button, ConfigProvider, Dropdown } from "antd"

const theme: ThemeConfig = {
    components: {
        Select: {
            colorBorder: "transparent",
            selectorBg: "transparent",
            colorIcon: "#fff",
            colorTextQuaternary: "#fff",
        },
    },
}

export default function Project() {
    const options: SelectProps["options"] = [
        { value: 123, label: "123" },
        { value: 123, label: "123" },
        { value: 123, label: "123" },
        { value: "new", label: "添加项目" },
    ]

    return (
        <ConfigProvider theme={theme}>
            <Dropdown
                className="w-full text-white"

            >
                <Button>
                    <span></span>
                    <IconTablerChevronDown className="text-lg" />
                </Button>
            </Dropdown>
        </ConfigProvider>
    )
}
