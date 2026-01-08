/// usr/bin/env jbang "$0" "$@" ; exit $?

//DEPS com.fasterxml.jackson.core:jackson-databind:2.19.2

import com.fasterxml.jackson.databind.ObjectMapper;

record Person(String name, int age) {
}

class App {
    public static void main(String[] args) throws Exception{
        var objectMapper = new ObjectMapper();

        var json = objectMapper.writeValueAsString(new Person("Julien", 37));

        System.out.println("json = " + json);

        var julien = objectMapper.readValue(json, Person.class);

        System.out.println("julien = " + julien);
    }

}
