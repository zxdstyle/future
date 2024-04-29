export default function Ethereum() {
    return (
        <div className="flex gap-2 rounded-3xl px-2 py-1 items-center bg-slate-100 ml-2 flex-shrink-0">
            <div className="flex items-center bg-white justify-center h-7 w-7 rounded-full">
                <IconFaBrandsEthereum className="text-gray-700" />
            </div>
            <span className="text-gray-700 text-sm font-bold">
                1,924
                <span> ETH</span>
            </span>
        </div>
    )
}
