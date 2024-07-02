import Manager from "./Manager"
import { Provider } from "./context"

interface Props {
}

export default function StorageManager({ }: Props) {
    return (
        <Provider>
            <Manager />
        </Provider>
    )
}
