import { Button, Popover } from "antd"

export default function Notification() {
    return (
        <Popover
            content={<Content />}
            placement="bottom"
            trigger="click"
            overlayInnerStyle={{ boxShadow: "none" }}
            overlayClassName="p-3 bg-white rounded-2xl shadow-normal"
        >
            <Button type="text" className="flex-shrink-0 w-8 px-0">
                <IconBiBell className="text-gray-700" />
            </Button>
        </Popover>
    )
}

function Content() {
    return (
        <div className="rounded-xl bg-white w-96">
            <div className="flex justify-between w-full mb-3">
                <div className="text-lg font-bold">Notifications</div>
                <div className="text-base font-medium cursor-pointer">Mark all read</div>
            </div>

            <div className="flex flex-col gap-y-4">
                <div className="rounded flex gap-3 cursor-pointer">
                    <div className="flex justify-center items-center rounded-xl w-12 h-12 flex-shrink-0 bg-gradient-to-l from-indigo-600 to-indigo-400">
                        <IconBiUpload className="text-white" />
                    </div>
                    <div className="flex flex-col">
                        <div className="mb-1 font-bold text-base">New Update: Horizon UI Dashboard PRO</div>
                        <div className="text-sm">A new update for your downloaded item is available!</div>
                    </div>
                </div>

                <div className="rounded flex gap-3 cursor-pointer">
                    <div className="flex justify-center items-center rounded-xl w-12 h-12 flex-shrink-0 bg-gradient-to-l from-indigo-600 to-indigo-400">
                        <IconBiUpload className="text-white" />
                    </div>
                    <div className="flex flex-col">
                        <div className="mb-1 font-bold text-base">New Update: Horizon UI Dashboard PRO</div>
                        <div className="text-sm">A new update for your downloaded item is available!</div>
                    </div>
                </div>
            </div>
        </div>
    )
}
