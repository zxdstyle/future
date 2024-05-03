import type { RouteObject } from "react-router-dom"

const projectRoutes: RouteObject[] = [
    { path: "/project", lazy: () => import("@/views/project") },
    { path: "/project/:id", lazy: () => import("@/views/project/detail") },
]

export default projectRoutes
