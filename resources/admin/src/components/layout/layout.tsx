import React from "react"
import { ThemedLayoutContextProvider } from "@refinedev/antd"
import { Layout as AntdLayout } from "antd"
import type { RefineThemedLayoutV2Props } from "@refinedev/antd"
import { Outlet } from "react-router-dom"
import DefaultSider from "./Sider"
import Navbar from "./Navbar"
import { NavbarProvider } from "@/components/layout/Navbar/context.tsx"

export const ThemedLayoutV2: React.FC<RefineThemedLayoutV2Props> = ({
    Sider,
    Title,
    Footer,
    OffLayoutArea,
    initialSiderCollapsed,
}) => {
    const SiderToRender = Sider ?? DefaultSider
    const hasSider = !!SiderToRender({ Title })

    return (
        <ThemedLayoutContextProvider initialSiderCollapsed={initialSiderCollapsed}>
            <AntdLayout className="h-full overflow-hidden" hasSider={hasSider}>
                {/* <SiderToRender Title={Title} /> */}
                <SiderToRender />

                <AntdLayout>
                    <NavbarProvider>
                        <Navbar />

                        <AntdLayout.Content>
                            <div className="overflow-hidden h-full">
                                <Outlet />
                            </div>

                            {OffLayoutArea && <OffLayoutArea />}
                        </AntdLayout.Content>

                        {Footer && <Footer />}
                    </NavbarProvider>
                </AntdLayout>
            </AntdLayout>
        </ThemedLayoutContextProvider>
    )
}
