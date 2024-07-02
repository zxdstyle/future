import type { PluginOption } from "vite"
import react from "@vitejs/plugin-react-swc"
import icon from "./icon"
import autoImport from "./autoImport"

/**
 * vite插件
 */
export function setupVitePlugins(): (PluginOption | PluginOption[])[] {
    const plugins: PluginOption[] = [react(), icon(), autoImport()]

    return plugins
}
