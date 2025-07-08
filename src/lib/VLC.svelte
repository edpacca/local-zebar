<script lang="ts">
    import { onMount } from "svelte";
    import type { VlcStatus } from "..";
    import SongInfo from "./SongInfo.svelte";

    const url = import.meta.env.VITE_VLC_PROXY_ENDPOINT;
    const defaultStatus: VlcStatus = {
        state: "",
        position: 0,
        meta: {}
    }
    let status: VlcStatus = defaultStatus;
    const getVlcStatus = async () => {
        try {
            const response = await fetch(
                url, {
                    method: "GET",
                    mode: "cors",
                    headers: {
                        "Accept": "application/json",
                        "Content-Type": "application/json",
                    }
                }
            )
            // if (!response.ok) {
            //     throw new Error(`HTTP Error: ${response.status}`);
            // }
            const data = await response.json();
            status = {
                state: data?.state ?? "",
                position: data?.position ?? 0,
                meta: { ...data?.information?.category?.meta ?? {}}
            }
            console.log(status);
        } catch (error) {
            status = defaultStatus;
            console.error("Error fetching VLC data: ", error);
        }
    }

    onMount(() => {
        getVlcStatus();
        setInterval(getVlcStatus, 5000);
    })
</script>

{#if status.state}
    <SongInfo artist={status.meta.artist} title={status.meta.title}/>
{:else}
    <div>novlc</div>
{/if}


