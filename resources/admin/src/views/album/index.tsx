import { useContext, useEffect } from "react"
import { Link } from "react-router-dom"
import CreateAlbum from "./create"
import { NavbarContext } from "@/components/layout"
import Album from "@/assets/images/album.png"

export function Component() {
    const { setLeftTools } = useContext(NavbarContext)
    useEffect(() => {
        setLeftTools(["navigator"])
    }, [])

    const items = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11]

    return (
        <div className="h-full p-12">
            <div className="h-full flex flex-wrap gap-10 justify-start content-start">

                <CreateAlbum />

                {items.map(item => (
                    <div className="relative text-center h-40 w-32 cursor-pointer group" key={item}>
                        <img className="absolute top-0 left-0 right-0 bottom-0 w-full h-full" src={Album} alt="" />
                        <h1 className="text-xl text-white relative pt-12">album</h1>
                        <div className="absolute top-0 w-full h-full bg-black/40 hidden group-hover:block transition-all duration-500">
                            <div className="w-full h-full flex flex-col items-center justify-end">
                                <div className="flex items-center justify-between py-3">
                                    <Link to={`/project/${item}`} className="text-white">
                                        <div className="flex items-center justify-center">
                                            <IconTablerEye />
                                            <span>预览</span>
                                        </div>
                                    </Link>
                                    <Link to={`/project/${item}`} className="text-white">
                                        <div className="flex items-center justify-center">
                                            <IconTablerEye />
                                            <span>详情</span>
                                        </div>
                                    </Link>
                                </div>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )
}
