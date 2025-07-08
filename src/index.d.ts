export interface VlcStatus {
    state: string,
    position: number,
    meta: VlcTrackMeta
}

export interface VlcTrackMeta {
    album?: string;
    filename?: string;
    artist?: string;
    title?: string;
}
