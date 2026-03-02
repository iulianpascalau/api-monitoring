export interface Metric {
    name: string;
    value: string;
    type: string;
    numAggregation: number;
    recordedAt: number;
    displayOrder: number;
    isAlarmEnabled?: boolean;
}

export interface MetricGroup {
    vmName: string;
    heartbeat: Metric | null;
    metrics: Metric[];
}
