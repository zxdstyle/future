import type { RouteObject } from "react-router-dom"

const storageRoutes: RouteObject[] = [
    { path: "/storage", lazy: () => import("@/views/storage") },
]

export default storageRoutes
