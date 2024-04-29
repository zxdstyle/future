import type { ThemeConfig } from "antd"
import { theme } from "antd"

const colors = {
    text: "#1B254B",
    primary: "#0088ff",
}

const customTheme: ThemeConfig = {
    algorithm: theme.defaultAlgorithm,
    hashed: false,
    token: {
        fontSize: 14,
        colorPrimary: colors.primary,
        colorTextBase: colors.text,
    },
    components: {
        Menu: {
            fontSize: 14,
        },
        Layout: {
            bodyBg: "#F8FAFC",
        },
        Form: {
            labelFontSize: 16,
        },
    },
}

export default customTheme

export { colors }
