<script>
    import {Route, useNavigate, useLocation} from 'svelte-navigator';
    import Login from './Login.svelte';
    import { user } from "./stores";

    const navigate = useNavigate();
    const location = useLocation();

    $: if (!$user) {
        console.log('nao autenticado')
        navigate("/login", {
            state: { from: $location.pathname },
            replace: true,
        });
    } else {
        console.log('AUTENTICADO')
    }
</script>

{#if $user}
    <h2 class="font-semibold">Olá {$user.username}</h2>
    <slot></slot>
{/if}

<Route path="login">
    <Login />
</Route>
