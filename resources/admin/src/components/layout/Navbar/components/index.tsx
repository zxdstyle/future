import Navigator from "./Navigator"
import QuickJumper from "./QuickJumper"

export const tools = {
    "navigator": <Navigator />,
    "quick-jumper": <QuickJumper />,
}

export type ToolKey = keyof typeof tools
