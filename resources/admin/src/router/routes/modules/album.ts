import type { RouteObject } from "react-router-dom"

const albumRoutes: RouteObject[] = [
    { path: "/album", lazy: () => import("@/views/album") },
    { path: "/album/:id", lazy: () => import("@/views/album/detail") },
]

export default albumRoutes
