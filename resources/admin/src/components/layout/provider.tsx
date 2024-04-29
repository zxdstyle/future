import { Outlet } from "react-router-dom"
import { useNotificationProvider } from "@refinedev/antd"
import routerBindings, { DocumentTitleHandler, UnsavedChangesNotifier } from "@refinedev/react-router-v6"
import { App as AntdApp } from "antd"
import { Refine } from "@refinedev/core"
// import { DevtoolsPanel, DevtoolsProvider } from "@refinedev/devtools"
// import { RefineKbar, RefineKbarProvider } from "@refinedev/kbar"
import { ColorModeContextProvider } from "@/contexts/color-mode"
import dataProvider from "@/services/provider"
import resources from "@/services/resources"

export function Provider() {
    return (
        <ColorModeContextProvider>
            <AntdApp className="h-full">
                {/* <DevtoolsProvider> */}
                <Refine
                    dataProvider={dataProvider}
                    notificationProvider={useNotificationProvider}
                    routerProvider={routerBindings}
                    resources={resources}
                    options={{
                        syncWithLocation: true,
                        warnWhenUnsavedChanges: true,
                        useNewQueryKeys: true,
                        projectId: "h1lfSS-pnet0i-tBr1R4",
                    }}
                >
                    {/* <RefineKbarProvider> */}
                    <Outlet />

                    {/* <RefineKbar /> */}

                    <UnsavedChangesNotifier />

                    <DocumentTitleHandler />
                    {/* </RefineKbarProvider> */}
                </Refine>
                {/* <DevtoolsPanel /> */}
                {/* </DevtoolsProvider> */}
            </AntdApp>
        </ColorModeContextProvider>
    )
}
