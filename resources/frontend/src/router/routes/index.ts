import { RouteDefinition } from '@solidjs/router';
import { lazy } from 'solid-js';
const routes: RouteDefinition[] = [
    {
        component: lazy(() => import('@/layouts/blank')),
        children: [
            { path: '/', component: lazy(() => import('@/views/home')) },
            { path: '/album', component: lazy(() => import('@/views/album')) },
            {
                path: '/album/:album_id',
                component: lazy(() => import('@/views/album/detail')),
            },
        ],
    },
    {
        component: lazy(() => import('@/layouts/admin')),
        children: [{ path: '/dashboard', component: lazy(() => import('@/views/admin/dashboard')) }],
    },
];

export default routes;
