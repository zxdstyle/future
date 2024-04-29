import { Router, useBeforeLeave } from '@solidjs/router';
import routes from '@/router/routes';

export default function RouterProvider() {
    return <Router>{routes}</Router>;
}
