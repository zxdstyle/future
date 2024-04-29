import React from "react"
import { useResource, useUserFriendlyName } from "@refinedev/core"
import { Layout as AntdLayout, Typography } from "antd"
import type { RefineThemedLayoutV2HeaderProps } from "@refinedev/antd"
import { Breadcrumb } from "@refinedev/antd"
import { Board, Ethereum, Notification, Profile, SearchBar } from "./components"

export const ThemedHeaderV2: React.FC<RefineThemedLayoutV2HeaderProps> = ({ sticky }) => {
    // const authProvider = useActiveAuthProvider()
    // const { data: user } = useGetIdentity({
    //     v3LegacyAuthProviderCompatible: Boolean(authProvider?.isLegacy),
    // })

    // const shouldRenderHeader = user && (user.name || user.avatar);
    //
    // if (!shouldRenderHeader) {
    //     return null;
    // }

    const { resource, identifier } = useResource()

    const headerStyles: React.CSSProperties = {
        backgroundColor: "transparent",
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        padding: "0px 32px",
        height: "76px",
        width: "100%",
    }

    if (sticky) {
        headerStyles.position = "sticky"
        headerStyles.top = 0
        headerStyles.zIndex = 1
    }
    const getUserFriendlyName = useUserFriendlyName()

    return (
        <AntdLayout.Header style={headerStyles}>
            <div className="flex flex-col justify-end h-full">
                <Breadcrumb breadcrumbProps={{ style: { fontSize: 12 } }} showHome />
                <Typography.Title level={3}>
                    {getUserFriendlyName(resource?.meta?.label ?? identifier, "plural")}
                </Typography.Title>
            </div>
            <div className="flex items-center bg-white px-3 py-2 rounded-full shadow-normal">
                <SearchBar />

                <Ethereum />

                <Notification />

                <Board />

                <Profile />
            </div>
        </AntdLayout.Header>
    )
}
