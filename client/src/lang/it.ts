const messages = {
    lang: 'Italiano',
    dashboard: {
        title: 'Dashboard',
        hello: 'Ciao { name }',
        latest_measurements: 'Ultime Misurazioni',
        sensors: 'Sensori',
    },
    login: {
        title: "Accesso"
    },
    sensors: {
        title: 'Sensori',
        create: {
            title: 'Crea Sensore',
            id_left_empty: "Questo sarà generato automaticamente se lasciato vuoto.",
            name: 'Nome',
            location: 'Posizione',
            type: 'Tipo',
            crt_button: 'Crea Sensore',
            threshold: 'Soglia',
            latitude: 'Latitudine',
            longitude: 'Longitudine',
            delete_button: 'Elimina Sensore',
        },
        list: {
            title: 'Lista Sensori',
            id: 'ID',
            name: 'Nome',
            type: 'Tipo',
            location: 'Posizione',
            accuracy: 'Accuratezza',
            accuracy_of: 'Accuratezza di { sensor }',
            value: 'Valore',
            sensor_name: 'Nome Sensore',
            timestamp: 'Data e ora',
            no_sensors: 'Nessun sensore disponibile',
            last_measurement: 'Ultima Misurazione',
            min: 'Min',
            max: 'Max',
            average: 'Media',
            status: 'Stato',
        },
        details: {
            title: 'Dettagli Sensore',
            id: 'ID',
            name: 'Nome',
            type: 'Tipo',
            location: 'Posizione',
        },
    },
    navbar: {
        choose_language: 'Scegli lingua',
    },
    logout: {
        title: 'Disconnetti',
    }
}

export default messages
