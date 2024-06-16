import Section from './section';
import { A } from '@solidjs/router';
import { SubtleButton } from '@/ui';
import { SeeMore } from '@/ui';
import { createSignal, For } from 'solid-js';
import Link from './link';
import clsx from 'clsx';

export default function () {
    const [locations] = createSignal([
        { id: 1, name: 'KeepSafe', color: '#D9188E' },
        { id: 2, name: 'Hidden', color: '#646278' },
        { id: 3, name: 'Projects', color: '#42D097' },
        { id: 4, name: 'Memes', color: '#A718D9' },
    ]);

    return (
        <Section
            name="Tags"
            actionArea={
                <A href="settings/library/locations">
                    <SubtleButton />
                </A>
            }
        >
            <SeeMore>
                {
                    <For each={locations()}>
                        {tag => (
                            <Link to={`location/${tag.id}`} class={clsx('border radix-state-open:border-accent', 'border-transparent')}>
                                <div class="h-[12px] w-[12px] shrink-0 rounded-full" style={{ 'background-color': tag.color || '#efefef' }} />
                                <span class="ml-1.5 truncate text-sm">{tag.name}</span>
                            </Link>
                        )}
                    </For>
                }
            </SeeMore>
        </Section>
    );
}
