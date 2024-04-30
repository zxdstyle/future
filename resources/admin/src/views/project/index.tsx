import { Image } from "antd"

export function Component() {
    return (
        <div className="p-12">
            <ul className="flex flex-wrap">
                <li>
                    <Image
                        width={64}
                        preview={false}
                        src="https://images.pexels.com/photos/9524657/pexels-photo-9524657.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2"
                    />
                </li>
            </ul>
        </div>
    )
}
