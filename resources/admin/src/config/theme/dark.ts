import type { ThemeConfig } from "antd"
import { theme } from "antd"

const colors = {
    base: "#1C1D26",
    overlay: "#16161d",
    text: "#ABACBA",
    primary: "#0088ff",
}

const customTheme: ThemeConfig = {
    algorithm: theme.darkAlgorithm,
    hashed: false,
    token: {
        fontSize: 14,
        colorPrimary: colors.primary,
        colorTextBase: colors.text,
        colorBgBase: colors.base,
        colorBgElevated: colors.overlay,
    },
    components: {
        Drawer: {
            colorPrimary: colors.base,
        },
        Menu: {
            fontSize: 14,
            itemBg: "transparent",
        },
        Layout: {
            bodyBg: colors.base,
            siderBg: colors.overlay,
        },
        Form: {
            labelFontSize: 16,
        },
    },
}

export default customTheme

export { colors }
