import type { MenuProps } from "antd"
import { Avatar, Menu, Popover } from "antd"
import { useNavigate } from "react-router-dom"
import { useGetIdentity } from "@refinedev/core"
// import { useAuthStore } from '@/store';

export default function Profile() {
    return (
        <Popover
            content={<Content />}
            placement="bottomRight"
            trigger="click"
            overlayInnerStyle={{ boxShadow: "none", padding: 0 }}
            overlayClassName="p-0 bg-white rounded-2xl shadow-normal overflow-hidden"
            className="ml-1"
            arrow
            destroyTooltipOnHide
        >
            <Avatar
                size="large"
                className="inline-flex items-center justify-center cursor-pointer flex-shrink-0"
                icon={<IconBxUser />}
            />
        </Popover>
    )
}

const items: MenuProps["items"] = [
    { key: "profile", label: "用户中心", icon: <IconBxUser /> },
    { key: "password", label: "修改密码", icon: <IconBxLock /> },
    { key: "logout", label: "退出登录", icon: <IconBxLogOut />, danger: true },
]

function Content() {
    const navigate = useNavigate()
    // const logout = useAuthStore(state => state.logout);
    // const user = useAuthStore(state => state.user);
    const { data: user } = useGetIdentity<IAdmin>()

    const handleSelect = ({ key }: { key: string }) => {
        switch (key) {
        case "logout":
            // logout();
            navigate("/login")
            return
        case "profile":
            navigate("/profile")
        }
    }

    return (
        <div className="font-medium w-40">
            <div className="border-b border-solid border-x-0 border-t-0 border-gray-200 px-5">
                <div className="py-3 w-full text-sm font-extrabold">
                    👋&nbsp; Hey,
                    {user?.name}
                </div>
            </div>
            <Menu items={items} onSelect={handleSelect} />
        </div>
    )
}
