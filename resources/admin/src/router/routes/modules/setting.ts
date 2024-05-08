import type { RouteObject } from "react-router-dom"

const settingRoutes: RouteObject[] = [
    {
        path: "/settings",
        lazy: () => import("@/views/settings"),
        children: [
            { path: "general", lazy: () => import("@/views/settings/general") },
            { path: "account", lazy: () => import("@/views/settings/account") },
        ],
    },
]

export default settingRoutes
