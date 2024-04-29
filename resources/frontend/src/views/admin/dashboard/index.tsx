import { createQuery } from '@tanstack/solid-query';
import { For, Match, Show, Switch } from 'solid-js';
import Folder from '@/assets/images/folder.png';
import Document from '@/assets/images/document.png';
import { A, useLocation, useNavigate } from '@solidjs/router';

export default function Dashboard() {
    const location = useLocation();
    const navigate = useNavigate();
    const state = createQuery<{ data: FileDescription[] }>(() => ({
        queryKey: ['files', location.query.pathname],
        queryFn: async () => {
            const dir = location.query.pathname ?? '/Users/zxdstyle';
            const resp = await fetch(`http://127.0.0.1:80/api/v1/files?dir=${dir}`);
            return resp.json();
        },
    }));

    const handleDbClick = (e: Event, file: FileDescription) => {
        e.stopPropagation();
        e.preventDefault();

        if (file.is_dir) {
            navigate(`/dashboard?pathname=${file.path}`);
        }
    };

    return (
        <div>
            <ul class="flex flex-wrap items-stretch gap-x-6 gap-y-4">
                <For each={state.data?.data}>
                    {item => (
                        <li
                            class="w-24 h-36 text-center p-2 rounded-xl active:bg-gray-100 hover:bg-gray-100"
                            onDblClick={e => handleDbClick(e, item)}
                        >
                            <Switch fallback={<img class="w-20 mx-auto" src={Document} alt={item.filename} />}>
                                <Match when={item.is_dir}>
                                    <img class="w-20 mx-auto" src={Folder} alt={item.filename} />
                                </Match>
                            </Switch>
                            <div class="text-sm break-words whitespace-normal line-clamp-2">{item.filename}</div>
                        </li>
                    )}
                </For>
            </ul>
        </div>
    );
}
