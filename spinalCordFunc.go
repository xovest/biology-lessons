package medicine.main

import (
    "fmt"
    "SpinalCordsFunctions"
    "time"
    "of"
    "info"
    )

func Reflex(by int ashColoredMatter) {
    var somatic.Reflex() += (receivers && effectors) in range(skin, ribbedMuscles);
    switch (somatic.Reflex()) {
        case 1: var monosynaptic.Reflex() <=> proprioceptive;
        case 2: var polisynaptic.Reflex() <=> exteroceptive;
    }
    
    monosynaptic.Reflex() {
        int receivers == new proprioReceivers[muscular];
        struct bow.Reflex() {
            int neuron := 2: {
                neuron.sensitive() => way(aferent);
                neuron.motor() => way(eferent);
            }
        }
        time.latence() == len(short);
        bool limited := true;
        bool iradiation := false;
        Role() {
            var tone[muscular] := adjusted();
            var position[body] := adjucted();
        }
    }
    
    polisynaptic.Reflex() {
        int receivers == range(skin) <=> new exteroReceivers;
        struct bow.Reflex() {
            int neuron := atleast(3): {
                type centers := polisynaptic;
                centers <- {neuron.sensitive(), neuron.association(), neuron.motor()};
            }
        }
        time.latence() == len(long); //because of neuron.intercalary() || neuron.association()
        bool iradiation := true;
        iradiation <-> intensity(stimulus); //according to Pfluger's laws
    }
    
    var vegetativeMedular.Reflex() += (receivers & effectors) in range(visceras, vessels);
    type.of vegetativeMedular.Reflex() {
        int pupilo+dilatator.Reflex();
        int cardio+accelerator.Reflex();
        int vaso+constrictoare.Reflex();
        int vaso+dilatatoare.Reflex();
        int sudoral.Reflex();
    }
    
    fmt.Println("End of Reflex Function");
}

func Leading(by int whiteMatter) {
    int leading.ways();
    switch (leading.ways()) {
        case 1: ways(ascending) <=> sensitive: lead(info) -> receivers;
                type.of ways(ascending) {
                    int ways(exteroceptive): lead(info) -> receivers.cutaneous();
                    int ways(proprioceptive): lead(info) -> proprioReceivers;
                    int ways(interoceptive): lead(info) -> visceras;
                }
        case 2: ways(motor): lead(commands);
                type.of ways(motor) {
                    int pyramidals;
                    int extraPyramidals;
                }
    }
    
    fmt.Println("End of Leading Function");
}

func main() {
    var SpinalCordsFunctions rune;
    _, err := fmt.Scanf(&SpinalCordsFunctions);
    fmt.Println(Reflex(SpinalCordsFunctions), Leading(SpinalCordsFunctions));
    return 0;
}