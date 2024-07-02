import { Button } from "antd"
import { useNavigate } from "react-router-dom"

export default function Navigator() {
    const navigate = useNavigate()

    return (
        <Button.Group>
            <Button onClick={() => navigate(-1)} type="text" icon={<IconPhArrowLeftBold />}></Button>
            <Button onClick={() => navigate(+1)} type="text" icon={<IconPhArrowLeftBold className="rotate-180" />}></Button>
        </Button.Group>
    )
}
