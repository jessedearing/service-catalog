insert into services (id, name, description) values
('5921e4a2-165b-4d11-87ce-22af741506ce', 'Locate Us', 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.'),
('c8a9a951-018a-41a2-8a2a-fe60dfe374bb', 'Collect Monday', ''),
('14f5a40b-81e0-48fd-b951-f82052ae6f40', 'Contact Us', 'Lorem ipsum dolor sit amet, consetetur sadipscing'),
('5a7a0edb-e627-4e1d-a9e4-1675c9906509', 'Contact Us', 'Lorem ipsum dolor sit amet, consetetur sadipscing'),
('a28509fc-cd84-4041-bb6b-83407a0ce3f9', 'FX Rates International', 'Lorem ipsum dolor'),
('f319bed6-6b4b-4f57-af53-c4da0962b197', 'FX Rates International', 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor'),
('fa38a530-c2d5-4739-9e1a-f2686656807e', 'Notifications', ''),
('8b1c562a-2ac8-41a4-944d-b3bd29953479', 'Notifications', ''),
('5e42ef5a-9fe2-46ba-9c33-325f9877b170', 'Priority Srevices', ''),
('fdf4d728-c056-447c-99aa-b7ca7a031303', 'Reporting', 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.'),
('1954568d-2d18-47e1-8bb0-1e0f4f99ac97', 'Security', 'Lorem ipsum dolor'),
('8eed1f42-71dc-47be-a470-582daa106bb5', 'Security', 'Lorem ipsum dolor');

insert into versions (id, service_id, version) values
(gen_random_uuid(), '5921e4a2-165b-4d11-87ce-22af741506ce', '1.0.0'),
(gen_random_uuid(), '5921e4a2-165b-4d11-87ce-22af741506ce', '1.1.0'),
(gen_random_uuid(), '5921e4a2-165b-4d11-87ce-22af741506ce', '1.1.1'),
(gen_random_uuid(), '5921e4a2-165b-4d11-87ce-22af741506ce', '1.2.0'),
(gen_random_uuid(), 'c8a9a951-018a-41a2-8a2a-fe60dfe374bb', '2.5.6'),
(gen_random_uuid(), 'c8a9a951-018a-41a2-8a2a-fe60dfe374bb', '2.5.7'),
(gen_random_uuid(), 'c8a9a951-018a-41a2-8a2a-fe60dfe374bb', '2.6.0'),
(gen_random_uuid(), 'c8a9a951-018a-41a2-8a2a-fe60dfe374bb', '2.6.1'),
(gen_random_uuid(), '14f5a40b-81e0-48fd-b951-f82052ae6f40', 'jaunty'),
(gen_random_uuid(), '14f5a40b-81e0-48fd-b951-f82052ae6f40', 'jezebel'),
(gen_random_uuid(), '14f5a40b-81e0-48fd-b951-f82052ae6f40', 'jovial'),
(gen_random_uuid(), '14f5a40b-81e0-48fd-b951-f82052ae6f40', 'jersey'),
(gen_random_uuid(), '5a7a0edb-e627-4e1d-a9e4-1675c9906509', '155.0.0'),
(gen_random_uuid(), '5a7a0edb-e627-4e1d-a9e4-1675c9906509', '156.0.0'),
(gen_random_uuid(), '5a7a0edb-e627-4e1d-a9e4-1675c9906509', '157.0.0'),
(gen_random_uuid(), '5a7a0edb-e627-4e1d-a9e4-1675c9906509', '158.0.0'),
(gen_random_uuid(), 'a28509fc-cd84-4041-bb6b-83407a0ce3f9', '1.0.0-alpha1'),
(gen_random_uuid(), 'a28509fc-cd84-4041-bb6b-83407a0ce3f9', '1.0.0-beta1'),
(gen_random_uuid(), 'a28509fc-cd84-4041-bb6b-83407a0ce3f9', '1.0.0-rc1'),
(gen_random_uuid(), 'a28509fc-cd84-4041-bb6b-83407a0ce3f9', '1.0.0'),
(gen_random_uuid(), 'f319bed6-6b4b-4f57-af53-c4da0962b197', '5.5.50'),
(gen_random_uuid(), 'f319bed6-6b4b-4f57-af53-c4da0962b197', '5.5.51'),
(gen_random_uuid(), 'f319bed6-6b4b-4f57-af53-c4da0962b197', '5.5.52'),
(gen_random_uuid(), 'f319bed6-6b4b-4f57-af53-c4da0962b197', '5.5.53'),
(gen_random_uuid(), 'fa38a530-c2d5-4739-9e1a-f2686656807e', '0.0.1025'),
(gen_random_uuid(), 'fa38a530-c2d5-4739-9e1a-f2686656807e', '0.0.1026'),
(gen_random_uuid(), 'fa38a530-c2d5-4739-9e1a-f2686656807e', '0.0.1027'),
(gen_random_uuid(), 'fa38a530-c2d5-4739-9e1a-f2686656807e', '0.0.1028'),
(gen_random_uuid(), '8b1c562a-2ac8-41a4-944d-b3bd29953479', '6e1663d78057b60e55407b21d9bfd7b5ebaa475e'),
(gen_random_uuid(), '8b1c562a-2ac8-41a4-944d-b3bd29953479', '272fb59d9c759a31da4fe7e5719f2aea5b17f959'),
(gen_random_uuid(), '8b1c562a-2ac8-41a4-944d-b3bd29953479', 'a83358ea7a95792c33ab83c54c4a271c223ed030'),
(gen_random_uuid(), '8b1c562a-2ac8-41a4-944d-b3bd29953479', '65b66010f55d72390a88d3af314168f7602bff8b'),
(gen_random_uuid(), '5e42ef5a-9fe2-46ba-9c33-325f9877b170', '1.0.0'),
(gen_random_uuid(), '5e42ef5a-9fe2-46ba-9c33-325f9877b170', '2.0.0'),
(gen_random_uuid(), '5e42ef5a-9fe2-46ba-9c33-325f9877b170', '4.0.0'),
(gen_random_uuid(), '5e42ef5a-9fe2-46ba-9c33-325f9877b170', '5.0.0'),
(gen_random_uuid(), 'fdf4d728-c056-447c-99aa-b7ca7a031303', 'v1.5.0'),
(gen_random_uuid(), 'fdf4d728-c056-447c-99aa-b7ca7a031303', 'v1.6.0'),
(gen_random_uuid(), 'fdf4d728-c056-447c-99aa-b7ca7a031303', 'v1.7.0'),
(gen_random_uuid(), 'fdf4d728-c056-447c-99aa-b7ca7a031303', 'v1.8.0'),
(gen_random_uuid(), '1954568d-2d18-47e1-8bb0-1e0f4f99ac97', 'v1'),
(gen_random_uuid(), '1954568d-2d18-47e1-8bb0-1e0f4f99ac97', 'v2'),
(gen_random_uuid(), '1954568d-2d18-47e1-8bb0-1e0f4f99ac97', 'v3'),
(gen_random_uuid(), '1954568d-2d18-47e1-8bb0-1e0f4f99ac97', 'v4'),
(gen_random_uuid(), '8eed1f42-71dc-47be-a470-582daa106bb5', '4.4.4'),
(gen_random_uuid(), '8eed1f42-71dc-47be-a470-582daa106bb5', '5.5.5'),
(gen_random_uuid(), '8eed1f42-71dc-47be-a470-582daa106bb5', '6.6.6'),
(gen_random_uuid(), '8eed1f42-71dc-47be-a470-582daa106bb5', '7.7.7');
