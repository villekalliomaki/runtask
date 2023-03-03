CREATE TABLE task_properties (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    created timestamp with time zone NOT NULL DEFAULT now(),
    description text
);
CREATE TABLE task_instance (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    created timestamp with time zone NOT NULL DEFAULT now(),
    started timestamp with time zone,

)