import type { RouteObject } from "react-router-dom"
import { handleModuleRoutes } from "../utils"
import { Provider, ThemedLayoutV2 } from "@/components/layout"
import ErrorComponent from "@/views/system/execption"

const modules = import.meta.glob("./modules/**/*.ts", { eager: true }) as Record<string, { default: RouteObject }>

const routes: RouteObject[] = [
    {
        path: "/",
        element: <Provider />,
        children: [
            {
                path: "/",
                element: <ThemedLayoutV2 />,
                children: [
                    {
                        path: "/",
                        errorElement: <ErrorComponent />,
                        children: handleModuleRoutes(modules),
                    },
                ],
            },
        ],
    },
]

export default routes
