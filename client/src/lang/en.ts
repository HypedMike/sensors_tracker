const messages = {
    lang: 'English',
    dashboard: {
        title: 'Dashboard',
        hello: 'Hello { name }',
        latest_measurements: 'Latest Measurements',
        sensors: 'Sensors',
    },
    login: {
        title: "Login"
    },
    sensors: {
        title: 'Sensors',
        create: {
            title: 'Create Sensor',
            id_left_empty: "This will be auto-generated if left empty.",
            name: 'Name',
            location: 'Location',
            type: 'Type',
            crt_button: 'Create Sensor',
            threshold: 'Threshold',
            latitude: 'Latitude',
            longitude: 'Longitude',
        },
        list: {
            title: 'Sensors List',
            id: 'ID',
            name: 'Name',
            type: 'Type',
            location: 'Location',
            accuracy: 'Accuracy',
            accuracy_of: 'Accuracy of { sensor }',
            value: 'Value',
            sensor_name: 'Sensor Name',
            timestamp: 'Timestamp',
            no_sensors: 'No sensors available',
            last_measurement: 'Last Measurement',
            min: 'Min',
            max: 'Max',
            average: 'Average',
            status: 'Status',
        },
        details: {
            title: 'Sensor Details',
            id: 'ID',
            name: 'Name',
            type: 'Type',
            location: 'Location',
        },
    },
    navbar: {
        choose_language: 'Choose language',
    },
    logout: {
        title: 'Logout',
    }
}
export default messages