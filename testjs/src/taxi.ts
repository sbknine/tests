export const Taxi = (distance: any, waitTime : any) => {
    const fare = 4*(Math.ceil(distance*2)/2) + Math.ceil(waitTime);
    return fare;
}

export const adjustFare = (f: number) => {
    return Math.max(f, 35);
}