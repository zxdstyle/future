import type { RouteObject } from "react-router-dom"

/**
 * 处理全部导入的路由模块
 * @param modules - 路由模块
 */
export function handleModuleRoutes(modules: Record<string, { default: RouteObject }>) {
    const routes: RouteObject[] = []

    Object.keys(modules).forEach((key) => {
        const item = modules[key].default
        if (item) {
            if (Array.isArray(item))
                routes.push(...item)
            else
                routes.push(item)
        }
        else {
            window.console.error(`路由模块解析出错: key = ${key}`)
        }
    })

    return routes
}
