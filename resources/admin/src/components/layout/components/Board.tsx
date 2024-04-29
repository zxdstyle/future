import { Button, Image, Popover } from "antd"

export default function Board() {
    return (
        <Popover
            content={<BoardContent />}
            placement="bottom"
            trigger="click"
            overlayInnerStyle={{ boxShadow: "none" }}
            overlayClassName="p-3 bg-white rounded-2xl shadow-normal"
        >
            <Button type="text" className="flex-shrink-0 w-8 px-0">
                <IconBxInfoCircle className="text-gray-700" />
            </Button>
        </Popover>
    )
}

function BoardContent() {
    return (
        <div>
            <Image className="rounded-2xl mb-5" preview={false} />
            <div className="flex flex-col">
                <Button type="primary" size="large" block>
                    Buy Horizon UI PRO
                </Button>
            </div>
        </div>
    )
}
